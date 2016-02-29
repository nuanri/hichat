//消息发送
function post_send_message() {
    var msg = $("#message-input")[0].value
    
    if (msg == "") {
        //alert("字符串为空")
    } else {
        $.ajax({
            type: "POST",
            url: "/api/messages",
            // The key needs to match your method's input parameter (case-sensitive).
            data: JSON.stringify({ "to": "Jian", "body": msg }),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function(data){
                if (data.error)
                    window.location = "/auth/signin"

                console.log("post message: data = " + data)
                //alert(data.status);
                //$("#send-body").append('<p>' + data.body + '</p>');
                $("#message-input")[0].value = ""
            },
            failure: function(errMsg) {
                alert(errMsg);
            }
        });
    }
}

//点击发送按钮发送
$( "#send-btn" ).click(function() {
    post_send_message()

});

//按 enter 键发送
$("#message-input").keypress(function(event) {
    if (event.which == 13) {
        event.preventDefault();
        post_send_message()
    }
});


// 获取新消息
function get_new_messages() {

    $.ajax({
        type: "GET",
        url: "/api/messages?t=" + new Date().toISOString(),
        // The key needs to match your method's input parameter (case-sensitive).
        //data: JSON.stringify({ "to": "Jian", "body": $("#message-input")[0].value }),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function(data){
            //alert(data.status);
           // console.log(data.signinuser)

            data.body.forEach(function(item){

               // if (item.username == "") {
                //window.location = "/auth/signin"
                //}
                //console.log(item.add_time)
                var ldate = new Date(item.add_time);
                var Y = ldate.getFullYear()
                var M = ldate.getMonth() + 1
                var D = ldate.getDate()
                var h = ldate.getHours()
                var m = ldate.getMinutes()
                var s = ldate.getSeconds()
                //ltime = Y+"年"+M+"月"+D+"日 "+h+":"+m+":"+s
                ltime = M+"月"+D+"日 "+h+":"+m

                if (item.username == data.signinuser){
                    $("#send-body").append('<div  class="show-message-me"><div class="date-me">'+ ltime+'</div><div class="show-username-me">' + item.username + "</div>" + ' <div class="msg-body"><div class="bubble-me me">' + item.msg + '</div></div></div>');
                } else {
                    $("#send-body").append('<div  class="show-message-you"><div class="date-you">'+ ltime+'</div><div class="show-username-you">' + item.username + "</div>" + ' <div class="msg-body"><div class="bubble-you you">' + item.msg + '</div></div></div>');
                }

                //在消息显示框内最底层，最后的输入总是显示在框内最后
                (function () {
                    var wtf = $('#send-body');
                    //返回整个元素的高度
                    var height = wtf[0].scrollHeight;
                    wtf.scrollTop(height);
                })();

            })

            $("#onlineuser").empty()
            data.onlineusers.forEach(function(username){
                $("#onlineuser").append( '<p> <span class="glyphicon glyphicon-user" aria-hidden="true"></span><span class="user-list">' + username + '</span></p>');
                (function () {
                    var wtf = $('#onlineuser');
                    var height = wtf[0].scrollHeight;
                    wtf.scrollTop(height);
                })();

            })

            setTimeout(function () {
                get_new_messages()
            }, 1000);

        },
        failure: function(errMsg) {
            alert(errMsg);
        }
    });

}


//auth_signup.html
//发送验证码
$( "#get-authcode-bt" ).bind("click", function() {
    $(this).unbind("click");

    $.ajax({
        url: "http://demo.hichat.xyz:8080/signup/request",
        method: "POST",
        // The key needs to match your method's input parameter (case-sensitive).
        data: JSON.stringify({"email": $("#inputEmail3")[0].value }),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function(data){
            console.log(data)
            $("#signup-error").empty()
            if (data.error)
                $("#signup-error").append(data.error)
            else 
                $("#authcode_key")[0].value = data.authcode_key

           // alert(data.status);
           //$("#message-input")[0].value = ""
        },
        failure: function(errMsg) {
            alert(errMsg);
        }
    });
})

$( "#signup-request-form" ).submit(function() {
  return false;
});


$("#signup-request-bt").click(function () {

    authcode = $("#authcode")[0].value
    authcode_key = $("#authcode_key")[0].value
    email = $("#inputEmail3")[0].value
    username = $("#username")[0].value
    password = $("#inputPassword3")[0].value

    $.ajax({
        url: "http://demo.hichat.xyz:8080/register/passwd",
        method: "POST",
        // The key needs to match your method's input parameter (case-sensitive).
        data: JSON.stringify({"authcode": authcode, "email":email, "authcode_key": authcode_key, "username": username, "password":password }),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function(data){
            console.log(data.authcode_key)
            $("#signup-error").empty()
            if (data.error) {
                $("#signup-error").append(data.error)
            }
            else {
                window.location = "/auth/signin"
            }

        },
        failure: function(errMsg) {
            alert(errMsg);
        }
    });
})

$("#signout").click(function (event) {
    event.preventDefault();
    $.ajax({
        url: $(this).attr('href'),
        cache: false,
        success: function(data){
            window.location = "/auth/signin"
        },
        failure: function(errMsg) {
            alert(errMsg);
        }
    });
})

$(document)
    .ready(
        function() {
            var h = ( $("html").height() - $("nav.navbar").height() - $("form").height() ) / 2;
            $("form").css("margin-top",h-30);
        }
    )
