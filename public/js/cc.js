
var o = JSON.parse(document.body.textContent);

console.log(o.events);

var e = o.events;
var insert = '';
for (const k in e) {
  var insert = insert + e[k]['segs'][0]['utf8'] + '<br>';
}
document.write(insert);
//console.log(insert);