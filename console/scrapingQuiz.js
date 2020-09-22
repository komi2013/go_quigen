
// var q = localStorage.q ? JSON.parse(localStorage.q) : [];
// var quiz = [];
// quiz[0] = $('h1').text();
// quiz[1] = '';
// $("#main-body > section:nth-child(1) > p").each(function( index ) {
//   console.log( $( this ).text() );
//   quiz[1] = quiz[1] + $( this ).text() + "\n";
// });
// q[q.length] = quiz;
// localStorage.q = JSON.stringify(q);
// var urls = localStorage.urls ? JSON.parse(localStorage.urls) : [];
// urls.shift();
// localStorage.urls = JSON.stringify(urls);
// setTimeout(function(){
//   location.href = urls[0];
// },1000);

// 1 ~500 e-words
// 500 ~ 1250 otsuka
// 
var q = [];
for (i = 7; i < 110; i++) {
  console.log( $("#post-940 > div > p:nth-child("+i+") > strong").text() );
  console.log( $("#post-940 > div > p:nth-child("+i+")").text() );
  var quiz = [$("#post-940 > div > p:nth-child("+i+") > strong").text(),
  $("#post-940 > div > p:nth-child("+i+")").text()];
  q[q.length] = quiz;
}
localStorage.q = JSON.stringify(q);
// console.log(JSON.stringify(q));
// #post-940 > div > p:nth-child(7) > strong
// #post-940 > div > p:nth-child(7)
// #post-940 > div > p:nth-child(8) > strong
// #post-940 > div > p:nth-child(86) > strong
// #post-940 > div > p:nth-child(109) > strong