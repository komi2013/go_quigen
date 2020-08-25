
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
  ['eel','https://translate.google.com/translate_tts?ie=UTF-8&q=eel&tl=en&total=1&idx=0&textlen=3&tk=598117.1042014&client=webapp']
  ,['Ray','https://translate.google.com/translate_tts?ie=UTF-8&q=Ray&tl=en&total=1&idx=0&textlen=3&tk=608318.1019397&client=webapp']
  ,['Tuna','https://translate.google.com/translate_tts?ie=UTF-8&q=Tuna&tl=en&total=1&idx=0&textlen=4&tk=504640.95611&client=webapp']
  ,['crab','https://translate.google.com/translate_tts?ie=UTF-8&q=crab&tl=en&total=1&idx=0&textlen=4&tk=961966.551829&client=webapp']
  ,['Panda','https://translate.google.com/translate_tts?ie=UTF-8&q=Panda&tl=en&total=1&idx=0&textlen=5&tk=759571.872744&client=webapp']
  ,['dugong','https://translate.google.com/translate_tts?ie=UTF-8&q=dugong&tl=en&total=1&idx=0&textlen=6&tk=18772.429935&client=webapp']
  ,['Wolf','https://translate.google.com/translate_tts?ie=UTF-8&q=Wolf&tl=en&total=1&idx=0&textlen=4&tk=915073.733370&client=webapp']
  ,['Lizard','https://translate.google.com/translate_tts?ie=UTF-8&q=Lizard&tl=en&total=1&idx=0&textlen=6&tk=826475.678480&client=webapp']
  ,['King cobra','https://translate.google.com/translate_tts?ie=UTF-8&q=King%20cobra&tl=en&total=1&idx=0&textlen=10&tk=300671.151620&client=webapp']	
  ,['chameleon','https://translate.google.com/translate_tts?ie=UTF-8&q=chameleon&tl=en&total=1&idx=0&textlen=9&tk=629728.1007067&client=webapp']
  ,['Ruffed lizard','https://translate.google.com/translate_tts?ie=UTF-8&q=Ruffed%20lizard&tl=en&tot…1&idx=0&textlen=13&tk=152344.300323&client=webapp']
];
for(var i = 0; i < audioArr.length; i++){
  var a = document.createElement("a");
  a.download = audioArr[i][0];
  a.href = audioArr[i][1];
  a.click();
}


