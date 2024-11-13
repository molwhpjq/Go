// $(document).ready(function(){
//     return {
//         // 验证码
//         inputCode: '', // 输入的值
//         checkCode: '', // 图片验证码的值
//         expressValue: '', // 表达式的值
//         // canvas各种设置
//         cvs: {
//             w: 300, // 给出默认宽度  宽度会在图片绘制时根据长度更改
//             h: 40, // 高 与input保持一致
//             fontSize: 24, // 字体大小
//             str: '+-*', // 符号生成范围
//             line: 3 // 噪音线个数
//         }
//     }
//
//
//
// // canvas
// // 随机整数生成器，范围[0, max)
//     $(document).ready(function(){
//     return Math.floor(Math.random() * 100000 % max);
// }
//
// // 生成随机表达式
// rCode() {
//     let a = this.rInt(100);
//     let b = this.rInt(10);
//     let op = this.cvs.str.charAt(this.rInt(this.cvs.str.length));
//     // 表达式
//     let code = `${a}${op}${b}=`;
//     this.checkCode = code;
//     // 表达式的值
//     // eslint-disable-next-line no-eval
//     this.expressValue = eval(code.substr(0, code.length - 1));
//     return code;
// }
//
// // 生成随机颜色 rgba格式
// rColor() {
//     let a = ((Math.random() * 5 + 5) / 10).toFixed(2);
//     return `rgba(${this.rInt(256)}, ${this.rInt(256)}, ${this.rInt(256)}, ${a})`;
// }
//
// // 验证码图片绘制
// drawCode(domCvs) {
//     console.log(domCvs);
//     let _this = this;
//     // 随机表达式
//     let checkCode = this.rCode();
//     // 宽设置
//     this.cvs.w = 10 + this.cvs.fontSize * this.checkCode.length;
//
//     // 判断是否支持canvas
//     if (domCvs !== null && domCvs.getContext && domCvs.getContext('2d')) {
//         // 设置显示区域大小
//         domCvs.style.width = _this.cvs.w;
//         console.log(_this.cvs.w);
//         // 设置画板宽高
//         domCvs.setAttribute('width', _this.cvs.w);
//         domCvs.setAttribute('height', 60);
//         // 画笔
//         let pen = domCvs.getContext('2d');
//         // 背景: 颜色  区域
//         pen.fillStyle = '#eee';
//         pen.fillRect(0, 10, _this.cvs.w, _this.cvs.h);
//         // 水平线位置
//         pen.textBaseline = 'middle'; // top middle bottom
//         // 内容
//         for (let i = 0; i < _this.checkCode.length; i ++) {
//             pen.fillStyle = _this.rColor(); // 随机颜色
//             pen.font = `bold ${_this.cvs.fontSize}px 微软雅黑`; // 字体设置
//             // 字符绘制: (字符, X坐标, Y坐标)
//             pen.fillText(checkCode.charAt(i), 10 + _this.cvs.fontSize * i, 30 + _this.rInt(10));
//         }
//         // 噪音线
//         for (let i = 0; i < _this.cvs.line; i ++) {
//             // 起点
//             pen.moveTo(_this.rInt(_this.cvs.w) / 2, _this.rInt(_this.cvs.h) + 10);
//             // 终点
//             pen.lineTo(_this.rInt(_this.cvs.w), _this.rInt(_this.cvs.h) + 10);
//             // 颜色
//             pen.strokeStyle = _this.rColor();
//             // 粗细
//             pen.lineWidth = '2';
//             // 绘制
//             pen.stroke();
//         }
//     } else {
//         this.$message.error('不支持验证码格式，请升级或更换浏览器重试');
//     }
// }
// // 点击时刷新二维码
// getCode() {
//     let domCvs = this.$refs.checkCode;
//     this.drawCode(domCvs);
// }
// // 点击登录时做的校验
// clickLogin() {
//     if (this.inputCode) {
//         // eslint-disable-next-line no-eval
//         if (eval(this.inputCode) === eval(this.expressValue)) {
//             // 验证成功要做的事
//             // this.$message.success('验证成功');
//         } else {
//             // 验证码有误
//             this.$message.error('验证码输入错误！');
//             return;
//         }
//     } else {
//         // 输入为空
//         this.$message.error('请输入右侧结果');
//         return;
//     }
//     //
//     let userCodeReg = /^[0-9a-zA-Z]+$/;
//     if (this.username === '') {
//         this.$message.error('请输入用户账号');
//         return;
//     }
//     if (! userCodeReg.test(this.username)) {
//         this.$message.error('用户账号不得使用中文或特殊符号');
//         return;
//     }
//     if (this.pwd === '') {
//         this.$message.error('请输入用户密码');
//         return;
//     }
//     this.login();
// }
// // 登录
// login() {
//     // 请求登录接口
// }
//
//     }