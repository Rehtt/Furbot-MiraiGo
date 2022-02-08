/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/8 15:54
 */

package models

type ApiResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Thumb string `json:"thumb"`
	Image []byte
}
