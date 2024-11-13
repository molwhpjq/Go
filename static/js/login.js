$(document).ready(function() {
    $('#btn').click(function() {
        var username = $('#username').val();
        var password = $('#password').val();
        var captcha = $('#captcha').val();
        if (username === '' || password === '') {
            $('#result').text('用户名或密码不能为空！');
        } else if (captcha !== '1234') {  // 这里假设正确的验证码是1234
            $('#result').text('验证码错误！');
        } else {
            // 如果所有验证都通过，这里可以执行一些登录操作，例如发送一个POST请求到服务器。
            $('#result').text('登录成功！');
        }
        alert(username)
        alert(password)
        $.ajax({
            url: '/dologin', // 后端控制器路由地址
            type: 'POST',
            dataType: 'json',
            contentType: 'application/json',
            data: JSON.stringify({uname: username, upwd: password}),
            success: function (response) {
                // 处理返回的JSON数据
                alert("返回的数据：" + response.message)
                console.log(response)
                if (response.success) {
                    $('#result').html('<p>登录成功！</p>');
                } else {
                    $('#result').html('<p>登录失败：' + response.message + '</p>');
                }
            },
            error: function (xhr, status, error) {
                console.log('Error occurred: ' + error);
            }
        });
    });

});

// $(document).ready(function() {
//     $('#btn').click(function() {
//         var username = $('#username').val();
//         var password = $('#password').val();
//         var captcha = $('#captcha').val();
//
//         // 假设这里是一些验证逻辑，例如检查用户名和密码是否为空，以及验证码是否正确。
//         if (username === '' || password === '') {
//             $('#result').text('用户名或密码不能为空！');
//         } else if (captcha !== '1234') {  // 这里假设正确的验证码是1234
//             $('#result').text('验证码错误！');
//         } else {
//             // 如果所有验证都通过，这里可以执行一些登录操作，例如发送一个POST请求到服务器。
//             $('#result').text('登录成功！');
//         }
//     });
// });

