// import { HandlerContext, Handlers } from "$fresh/server.ts";
// import axiod from "axiod";
// //op重定向接口
// export const handler: Handlers = {
//   GET(req: Request, ctx: HandlerContext) {
//     const url = new URL(req.url);
//     const params = url.searchParams;
//     const state = params.get("state");
//     const scope = params.get("scope");
//     const code = params.get("code");
//     const nonce = params.get("nonce");

//     axiod
//       .post(
//         "http://op.com:8000/v1/token",
//         //参数列表
//         {
//           grant_type: "authorization_code",
//           redirect_uri: "http://rp.com:8081/code_flow/oidc_op",
//           code: code,
//           nonce: nonce,
//           scope: scope,
//         },
//         //请求头配置  对client_id:client_secret进行base64标准填充编码
//         {
//           headers: { Authorization: "Basic cnAuY29tOnJwc2VjcmV0" },
//         }
//       )
//       .then((response) => {
//         //更新cookie
//         // 返回加密后的access_token 和 id_token
//         console.log(response.data);
//         //添加请求头信息
//         // console.log(response.data.Data);
//         axiod
//           .get("http://op.com:8000/v1/userinfo", {
//             //请求头发送access_token
//             headers: {
//               Authorization: `Bearer ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBjM01pT2lKdmNDNWpiMjBpTENKemRXSWlPaUl4SWl3aVlYVmtJam9pY25BdVkyOXRJaXdpWlhod0lqb3hOamszTlRNMk16QTVMQ0pwWVhRaU9qRTJPVGMxTXpRMU1Ea3NJbXAwYVNJNkltcDNkREF3TVNJc0ltNXZibU5sSWpvaUlpd2ljMk52Y0dVaU9pSWlmUS4zWVVVRjcyMEJaTHJfQWFLQWtpejlURng0N3NiNk10SmZaaV9TTnN4ckJz`,
//             },
//           })
//           .then((response) => {
//             //存储access_token
//             console.log(response.data);
//           });
//       });
//     const headers = new Headers();
//     //state
//     headers.set("location", "/");
//     return new Response(null, { status: 302, headers: headers });
//   },
// };

// // class message {
// //     code :number
// //     message :string
// //     data? :any
// //     constructor(code :number,message:string,data?:any) {
// //         this.code = code,
// //         this.message = message,
// //         this.data = data
// //     }
// // }
