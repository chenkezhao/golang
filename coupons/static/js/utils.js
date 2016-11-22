// $.fn.serializeObject = function() {
//     var o = {};
//     var a = this.serializeArray();
//     $.each(a, function() {
//         if (o[this.name]) {
//             if (!o[this.name].push) {
//                 o[this.name] = [o[this.name]];
//             }
//             o[this.name].push(this.value || '');
//         } else {
//             o[this.name] = this.value || '';
//         }
//     });
//     return o;
// };



//预览图片 
/*
	<div id="preview">
        <img id="imghead" src='/static/img/image.png' style="width: 50px;height: 50px">
    </div>
    <input id="prev" type="file" name="file" accept="image/jpeg,image/png" style="width: 190px;height: 30px;display: none;" onchange="doPreview(this,'prev','preview','imghead')" />
    <br/>
    <input type="button" onclick="$('#prev').click();" value="点击上传图片" style="margin-left:0px;margin-top:-5px; width: 120px;height: 30px;font-size: 16px;" />
    <br/>
    <br/>
    <span style="color:red">*图片大小不超过2MB</span>
    <br/>
    <span style="color:red">*缩略图片尺寸150*150</span>
*/
function doPreview(file,fileInputId,preViewId,preImgId) {
    var regext = /\.jpg$|\.jpeg$|\.png$/gi;
    var imgType = document.getElementById(fileInputId).value;
    // alert(imgPath);
    //判断图片类型
    if (!regext.test(imgType)) {
        errorMsg('对不起，系统仅支持 jpg | jpeg | png 格式的照片，请您调整格式后重新上传！');
        document.getElementById(fileInputId).value = "";
        return;
    }

    //文件大小不超过2MB
    //判断浏览器
    var isIE = /msie/i.test(navigator.userAgent) && !window.opera;
    var fileSize = 0;
    if (isIE && !file.files) {
        var filePath = file.value;
        var fileSystem = new ActiveXObject("Scripting.FileSystemObject");
        var files = fileSystem.GetFile(filePath);
        fileSize = files.Size;
    } else {
        fileSize = file.files[0].size;
    }
    console.log(fileSize);
    var size = fileSize / 1024;
    if (size > 2000) {
        errorMsg('图片大小超过2MB！');
        document.getElementById(fileInputId).value = "";
        return;
    }

    var div = document.getElementById(preViewId);
    if (file.files && file.files[0]) {
        //div.innerHTML = '<img id=preImgId title=点击查看大图  onclick=selectimg()>';
        var img = document.getElementById(preImgId);
        img.onload = function() {
            img.width = "150";
            img.height = "150";
        };
        var reader = new FileReader();
        reader.onload = function(evt) {
            img.src = evt.target.result;
            //img.title="双击查看大图";
        };
        reader.readAsDataURL(file.files[0]);
    } else {
        //var sFilter = 'filter:progid:DXImageTransform.Microsoft.AlphaImageLoader(sizingMethod=scale,src="';
        file.select();
        var src = document.selection.createRange().text;
        //div.innerHTML = '<img id=preImgId title=点击查看大图  onclick=selectimg()>';
        var img = document.getElementById(preImgId);
        img.filters.item('DXImageTransform.Microsoft.AlphaImageLoader').src = src;
        //img.filters.item('DXImageTransform.Microsoft.AlphaImageLoader').title="双击查看大图";
        //div.innerHTML = "<div id=divhead style='width:175px;height:175px;"+sFilter+src+"\"'></div>";
    }
}
