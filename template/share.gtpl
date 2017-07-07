<html>
<head>
    <meta http-equiv="Content-Type" content="charset=utf-8">
    <title>微信分享</title>
    <link rel="stylesheet" href="static/css/style.css">
</head>
<body>
微信分享<br>
<img src="http://demo.open.weixin.qq.com/jssdk/images/p2166127561.jpg" />
<div class="wxapi_container">
    <div class="lbox_close wxapi_form">
        <button class="btn btn_primary" id="onMenuShareAppMessage">分享到朋友圈</button>
    </div>
</div>
</body>
<script src="http://res.wx.qq.com/open/js/jweixin-1.2.0.js"></script>
<script>
  /*
   * 注意：
   * 1. 所有的JS接口只能在公众号绑定的域名下调用，公众号开发者需要先登录微信公众平台进入“公众号设置”的“功能设置”里填写“JS接口安全域名”。
   * 2. 如果发现在 Android 不能分享自定义内容，请到官网下载最新的包覆盖安装，Android 自定义分享接口需升级至 6.0.2.58 版本及以上。
   * 3. 常见问题及完整 JS-SDK 文档地址：http://mp.weixin.qq.com/wiki/7/aaa137b55fb2e0456bf8dd9148dd613f.html
   *
   * 开发中遇到问题详见文档“附录5-常见错误及解决办法”解决，如仍未能解决可通过以下渠道反馈：
   * 邮箱地址：weixin-open@qq.com
   * 邮件主题：【微信JS-SDK反馈】具体问题
   * 邮件内容说明：用简明的语言描述问题所在，并交代清楚遇到该问题的场景，可附上截屏图片，微信团队会尽快处理你的反馈。
   */
  wx.config({
      debug: true,
      appId: '{{.Appid}}',
      timestamp: [{{.Timestamp}}],
      nonceStr: '{{.Nonce}}',
      signature: '{{.Signature}}',
      jsApiList: [
        'checkJsApi',
        'onMenuShareTimeline',
        'onMenuShareAppMessage'
      ]
  });
</script>
<script>
wx.error(function (res) {
  alert(res.errMsg);
});
wx.ready(function() {
  document.querySelector('#onMenuShareAppMessage').onclick = function () {
    wx.onMenuShareAppMessage({
      title: '醉玲珑微信分享',
      desc: '醉玲珑微信分享到朋友圈',
      link: 'http://movie.douban.com/subject/25785114/',
      imgUrl: 'http://demo.open.weixin.qq.com/jssdk/images/p2166127561.jpg',
      trigger: function (res) {
      },
      success: function (res) {
        alert('分享成功');
      },
      cancel: function (res) {
        alert('分享取消');
      },
      fail: function (res) {
        alert(JSON.stringify(res));
      }
    });
  };

});

</script>
</html>