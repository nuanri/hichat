$( "#send-btn" ).click(function() {
    $.ajax({
        type: "POST",
        url: "http://192.168.0.7:8080/messages",
        // The key needs to match your method's input parameter (case-sensitive).
        data: JSON.stringify({ "to": "Jian", "body": $("#message-input")[0].value }),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function(data){
            //alert(data.status);
            //$("#send-body").append('<p>' + data.body + '</p>');
            $("#message-input")[0].value = ""
        },
        failure: function(errMsg) {
            alert(errMsg);
        }
    });

});

// 获取新消息
function get_new_messages() {

    $.ajax({
        type: "GET",
        url: "http://192.168.0.7:8080/messages?t=" + new Date().toISOString(),
        // The key needs to match your method's input parameter (case-sensitive).
        //data: JSON.stringify({ "to": "Jian", "body": $("#message-input")[0].value }),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function(data){
            //alert(data.status);
            data.body.forEach(function(item){
                $("#send-body").append('<p>' + item.add_time + ':' + item.msg + '</p>');
            })
            setTimeout(function () {
                get_new_messages()
            }, 5000);

        },
        failure: function(errMsg) {
            alert(errMsg);
        }
    });

}
