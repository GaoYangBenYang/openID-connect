// import { HandlerContext, Handlers } from "$fresh/server.ts";
// import axiod from "axiod";
// //op重定向接口
// export const handler: Handlers = {
//   GET(req: Request, ctx: HandlerContext) {
//     const client_secret: string = "rpsecret";
//     axiod
//       .post("http://op.com:8000/v1/registration", {
//         client_id: "rp.com",
//         client_secret: client_secret,
//       })
//       .then((response) => {
//         console.log(response.data);
//       });
//     return new Response("注册成功", { status: 200 });
//   },
// };
