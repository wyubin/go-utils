package vcf

import (
	"fmt"
	"strings"

	"ailab.com/vcfgo/utils/maptool"
)

// variant object for vcf row
type Variant struct {
	CHROM   string                 `json:"chrom"`
	POS     string                 `json:"pos"`
	ID      string                 `json:"id"`
	REF     string                 `json:"ref"`
	ALT     string                 `json:"alt"`
	QUAL    string                 `json:"qual"`
	FILTER  string                 `json:"filter"`
	INFO    map[string]interface{} `json:"info"`
	Samples []map[string]string    `json:"samples"`
}

// Unmarshal from string
func UnmarshalVariant(str string) (*Variant, error) {
	row := strings.Split(str, "\t")
	if len(row) < 8 {
		return nil, fmt.Errorf("invalid vcf row: %s", str)
	}
	variant := NewVariant(row[0], row[1], row[3], row[4])
	variant.ID = row[2]
	variant.QUAL = row[5]
	variant.FILTER = row[6]
	// INFO
	if row[7] != "." {
		for _, info := range strings.Split(row[7], ";") {
			kv := strings.SplitN(info, "=", 2)
			if len(kv) != 2 {
				continue
			}
			variant.INFO[kv[0]] = UnEscapeDataStr(kv[1])
		}
	}
	if len(row) < 9 {
		return variant, nil
	}
	// samples
	namesFORMAT := strings.Split(row[8], ":")
	for idx := 9; idx < len(row); idx++ {
		sample := map[string]string{}
		for k, v := range strings.Split(row[idx], ":") {
			if v == "." {
				continue
			}
			sample[namesFORMAT[k]] = UnEscapeDataStr(v)
		}
		variant.Samples = append(variant.Samples, sample)
	}
	return variant, nil
}

// Marshal to string
func (v *Variant) Marshal() string {
	strINFO := "."
	if len(v.INFO) > 0 {
		sliceINFO := []string{}
		for k, v := range v.INFO {
			sliceINFO = append(sliceINFO, fmt.Sprintf("%s=%s", k, EscapeDataStr(v.(string))))
		}
		strINFO = strings.Join(sliceINFO, ";")
	}
	res := []string{v.CHROM, v.POS, v.ID, v.REF, v.ALT, v.QUAL, v.FILTER, strINFO}
	if len(v.Samples) > 0 {
		aggregateFORMAT := map[string]string{}
		maptool.Update(aggregateFORMAT, v.Samples...)
		sliceFORMAT := maptool.Keys(aggregateFORMAT)
		strFORMAT := strings.Join(sliceFORMAT, ":")
		res = append(res, strFORMAT)
		for _, sample := range v.Samples {
			infosSample := []string{}
			for _, k := range sliceFORMAT {
				infoValue := "."
				if v, ok := sample[k]; ok {
					infoValue = v
				}
				infosSample = append(infosSample, infoValue)
			}
			res = append(res, strings.Join(infosSample, ":"))
		}
	}
	return strings.Join(res, "\t")
}
func NewVariant(chr, pos, ref, alt string) *Variant {
	return &Variant{
		CHROM:   chr,
		POS:     pos,
		ID:      ".",
		REF:     ref,
		ALT:     alt,
		QUAL:    ".",
		FILTER:  ".",
		INFO:    map[string]interface{}{},
		Samples: []map[string]string{},
	}
}
