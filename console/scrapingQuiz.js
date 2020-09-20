// $(function(){
  for (i = 0; i < 500; i++) {
    // console.log($('#ranking > table > tbody > tr:nth-child('+i+') > td:nth-child(2) > div > a').attr('href'));
    // console.log($('#ranking > table > tbody > tr:nth-child('+i+') > td:nth-child(2) > div > a').attr('href'));
    // setInterval(countup, 1000);
    var a = 1;
    setTimeout(function(){
      console.log($('#ranking > table > tbody > tr:nth-child('+a+') > td:nth-child(2) > div > a').attr('href'));
      a++;
    },i * 500);
  }
    // $("#ranking > table > tbody > tr").each(function( index ) {
    //   console.log( $( this ).children( "td:nth-child(2) > div > a") );
    // });  
  // });
  var urls = ["http://e-words.jp/w/%E3%82%B9%E3%82%AF%E3%83%AA%E3%83%97%E3%83%88.html"
  ,"http://e-words.jp/w/%E3%82%B3%E3%83%A1%E3%83%B3%E3%83%88.html"];
  for (i = 0; i < 1; i++) {
    console.log(urls[i]);
  }
  // function scraping() {
    var q = localStorage.q ? JSON.parse(localStorage.q) : [];
    // var numberI = localStorage.numberI ? localStorage.numberI : 0;
    var quiz = [];
    // console.log($('#word').text());
    // console.log($('#Summary').text());
    // #content > article > p:nth-child(6)
    quiz[0] = $('#word').text();
    quiz[1] = $('#Summary').text();
    // console.log(document.getElementsByXPath('//*[@id="content"]/article/p[3]')[0].firstChild.nodeValue);
    
    for (i = 1; i < $('#content > article > p').length+1; i++) {
      // console.log($('#content > article > p:nth-child('+i+')').firstChild.nodeValue);
      quiz[1] = quiz[1] + document.getElementsByXPath('//*[@id="content"]/article/p['+i+']')[0].innerHTML;
      // console.log(document.getElementsByXPath('//*[@id="content"]/article/p['+i+']')[0].innerHTML);
    }
    q[q.length] = quiz;
    // console.log(quiz);
    // JSON.parse(localStorage.quiz);
    localStorage.q = JSON.stringify(q);
    localStorage.numberI++;
    if(localStorage.numberI < urls.length){
      location.href = urls[localStorage.numberI];
    }
    
  // }
  