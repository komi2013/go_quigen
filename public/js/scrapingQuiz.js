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
arr.push(document.getElementsByXPath('//*[@id="mainCol"]/div[2]/div[1]')[0].firstChild.nodeValue);

// console.log($('#select_a').html());
if ( $('#select_a').html() ) {
  arr.push($('#select_a').html());
  arr.push($('#select_i').html());
  arr.push($('#select_u').html());
  arr.push($('#select_e').html());
} else {
  arr.push("");
  arr.push("");
  arr.push("");
  arr.push("");
}

arr.push($('#answerChar').html());
// console.log($('#kaisetsu div').eq(-0).html());
if ($('#kaisetsu div') && $('#kaisetsu div').eq(-0)) {
  // console.log(document.getElementsByXPath('//*[@id="kaisetsu"]/div[1]')[0].firstChild.nodeValue);
  arr.push($('#kaisetsu div').eq(-0).html());

}


var localQ = localStorage.q ? JSON.parse(localStorage.q) : [];
localQ.unshift(arr);
localStorage.q = JSON.stringify(localQ);

// $('.submit').click();

let csvContent = "data:text/csv;charset=utf-8,";
localQ.forEach(function(rowArray) {
  let row = rowArray.join(",");
  csvContent += row + "\r\n";
});

// console.log(arr);
// console.log(csvContent);
var encodedUri = encodeURI(csvContent);
window.open(encodedUri);
// console.log(document.getElementsByXPath('//*[@id="mainCol"]/div[2]/div[1]/div/img')[0].src);

// 平成25年春期　問1
// 平成31年春期　問3
