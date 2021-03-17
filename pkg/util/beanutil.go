package util

import "encoding/json"

/**

 */
func BeanCopy(dst, src interface{}) error {
	marshal, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(marshal, dst)
	return err
}
