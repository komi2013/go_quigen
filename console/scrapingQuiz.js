
var q = localStorage.q ? JSON.parse(localStorage.q) : [];
var quiz = [];
quiz[0] = $('h1').text();
quiz[1] = '';
$("#main-body > section:nth-child(1) > p").each(function( index ) {
  console.log( $( this ).text() );
  quiz[1] = quiz[1] + $( this ).text() + "\n";
});
q[q.length] = quiz;
localStorage.q = JSON.stringify(q);
var urls = localStorage.urls ? JSON.parse(localStorage.urls) : [];
urls.shift();
localStorage.urls = JSON.stringify(urls);
setTimeout(function(){
  location.href = urls[0];
},1000);

  