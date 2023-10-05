package test

import (
	"OpenIDProvider/internal/utils"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//编码

	// jwt := model.NewJWT(model.NewHeader("HS256", "JWT"), model.NewPayload("op.com", "5003", "rp.com", "oyvcp3q91sohn3zfebosc", ""))

	// fmt.Println(utils.EncodeTheJWT(jwt))
	fmt.Println(utils.DecodeTheJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJvcC5jb20iLCJzdWIiOiI1MDAzIiwiYXVkIjoicnAuY29tIiwibmJmIjoxNjk2NDkyMTM0LCJleHAiOjE2OTY0OTM5MzQsImlhdCI6MTY5NjQ5MjEzNCwianRpIjoib3l2Y3AzcTkxc29objN6ZmVib3NjIiwic2Vzc2lvbl9zdGF0ZSI6IiJ9.6Z3jP5F_qQ2uHXQP9SbdsHVeag6JH83r6Cp67siM280"))
	fmt.Println(utils.Base64RawURLEncoding("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJvcC5jb20iLCJzdWIiOiI1MDAzIiwiYXVkIjoicnAuY29tIiwibmJmIjoxNjk2NDkyMTM0LCJleHAiOjE2OTY0OTM5MzQsImlhdCI6MTY5NjQ5MjEzNCwianRpIjoib3l2Y3AzcTkxc29objN6ZmVib3NjIiwic2Vzc2lvbl9zdGF0ZSI6IiJ9.6Z3jP5F_qQ2uHXQP9SbdsHVeag6JH83r6Cp67siM280"))

}
