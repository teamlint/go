package templatex

import "errors"

// dict 字典参数方法
func dict(values ...interface{}) (map[string]interface{}, error) {
	defer recovery()
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
func set(d map[string]interface{}, key string, value interface{}) map[string]interface{} {
	d[key] = value
	return d
}

func unset(d map[string]interface{}, key string) map[string]interface{} {
	delete(d, key)
	return d
}

func hasKey(d map[string]interface{}, key string) bool {
	_, ok := d[key]
	return ok
}

func pluck(key string, d ...map[string]interface{}) []interface{} {
	res := []interface{}{}
	for _, dict := range d {
		if val, ok := dict[key]; ok {
			res = append(res, val)
		}
	}
	return res
}

func keys(dicts ...map[string]interface{}) []string {
	k := []string{}
	for _, dict := range dicts {
		for key := range dict {
			k = append(k, key)
		}
	}
	return k
}

func pick(dict map[string]interface{}, keys ...string) map[string]interface{} {
	res := map[string]interface{}{}
	for _, k := range keys {
		if v, ok := dict[k]; ok {
			res[k] = v
		}
	}
	return res
}

func omit(dict map[string]interface{}, keys ...string) map[string]interface{} {
	res := map[string]interface{}{}

	omit := make(map[string]bool, len(keys))
	for _, k := range keys {
		omit[k] = true
	}

	for k, v := range dict {
		if _, ok := omit[k]; !ok {
			res[k] = v
		}
	}
	return res
}

// func merge(dst map[string]interface{}, srcs ...map[string]interface{}) interface{} {
// 	for _, src := range srcs {
// 		if err := mergo.Merge(&dst, src); err != nil {
// 			// Swallow errors inside of a template.
// 			return ""
// 		}
// 	}
// 	return dst
// }

func values(dict map[string]interface{}) []interface{} {
	values := []interface{}{}
	for _, value := range dict {
		values = append(values, value)
	}

	return values
}
