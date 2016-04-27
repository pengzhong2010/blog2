$(function(){
    $(".product_edit").click(function(){
        choose_arr=new Array();
        $(".choose:checked").each(function(){
            tr=$(this).parent().parent().parent().parent();
            choose_arr.push($(".pid",tr).html());
        });
        if(choose_arr.length!=1){
            alert("请选择正确的条目数");
        }else{
            window.location.href="/product/edit/"+choose_arr[0];
        }
    });

    $(".product_delete").click(function(){
        choose={};
        choose.data= new Array();
        $(".choose:checked").each(function(){
            tr=$(this).parent().parent().parent().parent();
            choose.data.push($(".pid",tr).html());
        });
        if(choose.data.length==0){
            alert("请选择正确的条目数");
        }else{
            f=confirm('确定删除？');
            if(!f) return false;
            str=JSON.stringify(choose);
            str_display='';
            $.post('/product/delete/',{pids:str,"csrfmiddlewaretoken":s('csrftoken')},function(data){
                console.log(data);
                for(i in data){
                    console.log(i);
                    str_display+='ID:  '+i+'    STATE:  ';
                    str_display+=(data[i]==1)?'Success':'False';
                    str_display+="\n";
                }
                alert(str_display);
                location.reload(true);
            },"json");
        }
    });

    $(".page_form").submit(function(){
        page=$("input[name='p']",$(this)).val();
//        console.log(page);
        if(!page)
        {
            alert("请填写访问页号");
            return false;
        }

    });

    $(".search_form").submit(function(){
        page=$("input[name='search']",$(this)).val();
//        console.log(page);
        if(!page)
        {
            alert("请填写搜索字段");
            return false;
        }

    });

});