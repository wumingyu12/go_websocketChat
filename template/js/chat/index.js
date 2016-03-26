
var host = '127.0.0.1';    
var port = 8080;    
var url = 'ws://'+host+':'+port+'/';    
var w = new WebSocket(url);  //构建于服务器之间的连接通信  
<!--var audioElement = document.createElement('audio');  -->  
   
<!--audioElement.setAttribute('src', 'qqmsg.mp3');-->    

w.onopen = function()//通过onopen事件句柄来监听socket的打开事件  
{    
     $('chat-box').innerHTML = '已连接到服务器......<br/>';    
}

w.onmessage = function(e)//用onmessage事件句柄接受服务器传过来的数据  
{    
   var msg = e.data;    
   var chatBox = $('chat-box');    
   // audioElement.play();        
   chatBox.innerHTML = chatBox.innerHTML+msg+'<br/>';    
}    
       
function send()//使用send方法向服务器发送数据  
{    
     var talk = $('talk');    
     var nike = $('nike');     
     w.send('<strong style="color:red">'+nike.value+':</strong>'+talk.value);    
}    
 
 function $(id)  
 {    
     return document.getElementById(id);    
 }    