package vcf

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/wyubin/go-utils/maptool"
)

var (
	GeneralTagTypes = []string{"INFO", "FORMAT", "contig", "SAMPLE", "FILTER"}
	HeaderVcfFormat = NewTag("attr", "fileformat", map[string]string{"Description": "VCFv4.3"})
	VCF_COLNAMES    = []string{"CHROM", "POS", "ID", "REF", "ALT", "QUAL", "FILTER", "INFO"}
	AttrOrder       = []string{"Number", "Type"}
)

type Header struct {
	ID   string
	Tag  string // INFO or FORMAT or contig or attr
	Args map[string]string
}

func (s *Header) String() string {
	var strInfo string
	switch s.Tag {
	case "INFO", "FORMAT", "contig", "SAMPLE", "FILTER":
		infoPairs := []string{fmt.Sprintf("ID=%s", s.ID)}
		if len(s.Args) > 0 {
			infoPairs = append(infoPairs, s.getArgsString())
		}
		strInfo = fmt.Sprintf("##%s=<%s>", s.Tag, strings.Join(infoPairs, ","))
	case "attr":
		strInfo = fmt.Sprintf("##%s=%s", s.ID, s.Args["Description"])
	default:
		strInfo = ""
	}
	return strInfo
}

func (s *Header) getArgsString() string {
	substrs := []string{}
	for _, k := range AttrOrder {
		v, ok := s.Args[k]
		if ok {
			substrs = append(substrs, fmt.Sprintf("%s=%s", k, v))
		}
	}
	for k, v := range s.Args {
		if slices.Contains(AttrOrder, k) {
			continue
		}
		substrs = append(substrs, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(substrs, ",")
}
func ParseHeader(vcfInput io.Reader) ([]Header, error) {
	res := []Header{}
	resErr := []error{}
	var line string
	scanner := bufio.NewScanner(vcfInput)
	for scanner.Scan() {
		// Get first line to check chr and exit
		line = scanner.Text()
		if !strings.HasPrefix(line, "#") {
			break
		} else if !strings.HasPrefix(line, "##") {
			// one # is colnames, try ti get SAMPLE
			cols := strings.Split(line, "\t")
			for i, col := range cols {
				if i < 9 {
					continue
				}
				headerObj := NewTag("SAMPLE", col, map[string]string{"Description": col})
				res = append(res, headerObj)
			}
		} else {
			headerTmp, err := HeaderValueUnmarshal(line)
			resErr = append(resErr, err)
			res = append(res, *headerTmp)
		}
	}
	return res, errors.Join(resErr...)
}

func DumpHeader(fileHeader io.Writer, headers []Header) error {
	nameSamples := []string{}
	for _, header := range headers {
		fileHeader.Write([]byte(header.String() + "\n"))
		if header.Tag == "SAMPLE" {
			nameSamples = append(nameSamples, header.ID)
		}
	}
	colnames := VCF_COLNAMES[:]
	if len(nameSamples) > 0 {
		colnames = append(colnames, append([]string{"FORMAT"}, nameSamples...)...)
	}
	strCols := "#" + strings.Join(colnames, "\t") + "\n"
	fileHeader.Write([]byte(strCols))
	return nil
}

// transform header string into struct
func HeaderValueUnmarshal(headerStr string) (*Header, error) {
	var header Header
	strTag := strings.TrimSpace(strings.TrimPrefix(headerStr, "##"))
	headerTuple := strings.SplitN(strTag, "=", 2)
	if len(headerTuple) < 2 {
		return nil, fmt.Errorf("tag[%s] not valid HeaderValue", headerStr)
	}
	// parse headerStr
	if !strings.HasPrefix(headerTuple[1], "<") || !strings.HasSuffix(headerTuple[1], ">") {
		// attrs
		header = NewTag("attr", headerTuple[0], map[string]string{"Description": headerTuple[1]})
		return &header, nil
	}
	headerStrClear := strings.Trim(headerTuple[1], "<> ")
	idxStart := 0
	idxLast := len(headerStrClear)
	headerMap := map[string]string{}
	for idxStart < idxLast {
		regionCurr := headerStrClear[idxStart:]
		idxNameEnd := strings.Index(regionCurr, "=")
		_key := regionCurr[:idxNameEnd]
		var _val string
		if strings.HasPrefix(regionCurr[idxNameEnd+1:], "\"") {
			idxValueEnd := strings.Index(regionCurr[idxNameEnd+2:], "\"") + idxNameEnd + 2
			_val = strconv.Quote(regionCurr[idxNameEnd+2 : idxValueEnd])
			idxStart += idxValueEnd + 2
		} else {
			idxValueEndTmp := strings.Index(regionCurr[idxNameEnd+1:], ",")
			var idxValueEnd int
			if idxValueEndTmp == -1 {
				idxValueEnd = len(regionCurr)
			} else {
				idxValueEnd = idxValueEndTmp + idxNameEnd + 1
			}
			_val = regionCurr[idxNameEnd+1 : idxValueEnd]
			idxStart += idxValueEnd + 1
		}
		headerMap[_key] = _val
	}
	id, _ := maptool.Pop[string, string](headerMap, "ID")
	header = NewTag(headerTuple[0], id, headerMap)
	return &header, nil
}

func NewTag(typeTag, id string, attrs ...map[string]string) Header {
	header := Header{
		ID:   id,
		Tag:  typeTag,
		Args: map[string]string{},
	}
	argsDefault := map[string]string{}
	switch typeTag {
	case "INFO", "FORMAT":
		argsDefault = map[string]string{"Number": ".", "Type": "String", "Description": strconv.Quote(id)}
	case "FILTER", "SAMPLE", "attr":
		argsDefault = map[string]string{"Description": ""}
	}
	extendArgs := append([]map[string]string{argsDefault}, attrs...)
	for _, arg := range extendArgs {
		maptool.Update(header.Args, arg)
	}
	return header
}

func FuncContains(infoTag string) func(Header) bool {
	return func(header Header) bool {
		return header.ID == infoTag
	}
}
