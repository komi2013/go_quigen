
var o = JSON.parse(document.body.textContent);

console.log(o.events);

var e = o.events;
var insert = '';
for (const k in e) {
  var insert = insert + e[k]['segs'][0]['utf8'] + '<br>';
}
document.write(insert);
//console.log(insert);

var audioArr = [

];

var i2 = 0;
for(var i = 0; i < audioArr.length; i++){
  setTimeout(function() {
    var a = document.createElement("a");

    a.download = audioArr[i2][0];
    a.href = audioArr[i2][1];
    a.click();
    console.log(audioArr[i2][0]);
    console.log(audioArr[i2][1]);
    i2++;
  }, i * 5000);
}


