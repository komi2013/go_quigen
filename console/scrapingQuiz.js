document.getElementsByXPath = function(expression, parentElement) {
  var r = []
  var x = document.evaluate(expression, parentElement || document, null, XPathResult.ORDERED_NODE_SNAPSHOT_TYPE, null)
  for (var i = 0, l = x.snapshotLength; i < l; i++) {
      r.push(x.snapshotItem(i))
  }
  return r
}
// console.log(document.getElementsByXPath('//*[@id="mainCol"]/div[2]/div[1]')[0].firstChild.nodeValue);

var arr = [];

// #mainCol > div.main.kako.doujou > div:nth-child(4)
arr.push($('#mainCol > div.main.kako.doujou > div:nth-child(4)').html());

// console.log($('#select_a').html());
var choice_0 = '';
if ( $('#select_a').html() ) {
  if ( $('#answerChar').html() == 'ア' ) {
    choice_0 = $('#select_a').html()
    $('.selectBtn')[0].click();
  }else{
    arr.push($('#select_a').html());
  }
  if ( $('#answerChar').html() == 'イ' ) {
    choice_0 = $('#select_i').html()
    $('.selectBtn')[1].click();
  }else{
    arr.push($('#select_i').html());
  }
  if ( $('#answerChar').html() == 'ウ' ) {
    choice_0 = $('#select_u').html()
    $('.selectBtn')[2].click();
  }else{
    arr.push($('#select_u').html());
  }
  if ( $('#answerChar').html() == 'エ' ) {
    choice_0 = $('#select_e').html()
    $('.selectBtn')[3].click();
  }else{
    arr.push($('#select_e').html());
  }
} else {
  if ( $('#answerChar').html() == 'ア' ) {
    choice_0 = 'ア'
  }else{
    arr.push('ア');
  }
  if ( $('#answerChar').html() == 'イ' ) {
    choice_0 = 'イ'
  }else{
    arr.push('イ');
  }
  if ( $('#answerChar').html() == 'ウ' ) {
    choice_0 = 'ウ'
  }else{
    arr.push('ウ');
  }
  if ( $('#answerChar').html() == 'エ' ) {
    choice_0 = 'エ'
  }else{
    arr.push('エ');
  }
}
arr.push(choice_0);
// arr.push($('#answerChar').html());
// console.log($('#kaisetsu div').eq(-0).html());
if ($('#kaisetsu div') && $('#kaisetsu div').eq(-0)) {
  // console.log(document.getElementsByXPath('//*[@id="kaisetsu"]/div[1]')[0].firstChild.nodeValue);
  arr.push($('#kaisetsu div').eq(-0).html());

}

if( $('#answerChar').html() ){
  var localQ = localStorage.q ? JSON.parse(localStorage.q) : [];
  localQ.unshift(arr);
  localStorage.q = JSON.stringify(localQ);
  $('.submit').click();
}


// let csvContent = "data:text/csv;charset=utf-8,";
// localQ.forEach(function(rowArray) {
//   let row = rowArray.join(",");
//   csvContent += row + "\r\n";
// });

// console.log(arr);
// console.log(csvContent);
// var encodedUri = encodeURI(csvContent);
// window.open(encodedUri);
// console.log(document.getElementsByXPath('//*[@id="mainCol"]/div[2]/div[1]/div/img')[0].src);

// 平成25年春期　問1
// 平成31年春期　問3
