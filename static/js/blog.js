//登录
$("#login-form").validate({
    rules:{
        username:{
            required:true,
            rangelength:[5,10]
        },
        password:{
            required:true,
            rangelength:[5,10]
        }
    },
    messages:{
        username:{
            required:"请输入用户名",
            rangelength:"用户名必须是5-10位"
        },
        password:{
            required:"请输入密码",
            rangelength:"密码必须是5-10位"
        }
    },
    submitHandler:function (form) {
        var urlStr ="/login"
        alert("urlStr:"+urlStr)
        $(form).ajaxSubmit({
            url:urlStr,
            type:"post",
            dataType:"json",
            success:function (data,status) {
                alert("data:"+data.message+":"+status)
                if(data.code == 1){
                    setTimeout(function () {
                        window.location.href="/"
                    },1000)
                }
            },
            error:function (data,status) {
                alert("err:"+data.message+":"+status)
            }
        });
    }
});