package gopium

import "encoding/json"

func remarshal(in, out any) error {
	data, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, out)
}

func splitRunes(value string) []string {
	if value == "" {
		return nil
	}

	out := make([]string, 0, len(value))
	for _, r := range value {
		out = append(out, string(r))
	}
	return out
}
