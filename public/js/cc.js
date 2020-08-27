
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
   ['Capybara','https://translate.google.com/translate_tts?ie=UTF-8&q=Capybara&tl=en&total=1&idx=0&textlen=8&tk=526818.970630&client=webapp ']
  ,['Kamileka','https://translate.google.com/translate_tts?ie=UTF-8&q=Kamileka&tl=en&total=1&idx=0&textlen=8&tk=589208.932860&client=webapp ']
  ,['platypus','https://translate.google.com/translate_tts?ie=UTF-8&q=platypus&tl=en&total=1&idx=0&textlen=8&tk=170706.285878&client=webapp ']
  ,['fur seal','https://translate.google.com/translate_tts?ie=UTF-8&q=fur%20seal&tl=en&total=1&idx=0&textlen=8&tk=9079.451859&client=webapp ']
  ,['Fox','https://translate.google.com/translate_tts?ie=UTF-8&q=Fox&tl=en&total=1&idx=0&textlen=3&tk=22335.430427&client=webapp ']
  ,['Squirrel','https://translate.google.com/translate_tts?ie=UTF-8&q=Squirrel&tl=en&total=1&idx=0&textlen=8&tk=911747.733159&client=webapp ']
  ,['Rhino','https://translate.google.com/translate_tts?ie=UTF-8&q=Rhino&tl=en&total=1&idx=0&textlen=5&tk=194801.278165&client=webapp ']
  ,['zebra','https://translate.google.com/translate_tts?ie=UTF-8&q=zebra&tl=en&total=1&idx=0&textlen=5&tk=172614.287778&client=webapp ']
  ,['Grizzly','https://translate.google.com/translate_tts?ie=UTF-8&q=Grizzly&tl=en&total=1&idx=0&textlen=7&tk=436675.28583&client=webapp ']
  ,['Koala','https://translate.google.com/translate_tts?ie=UTF-8&q=Koala&tl=en&total=1&idx=0&textlen=5&tk=68929.510757&client=webapp ']
  ,['Bat','https://translate.google.com/translate_tts?ie=UTF-8&q=Bat&tl=en&total=1&idx=0&textlen=3&tk=770803.853143&client=webapp ']
  ,['armadillo','https://translate.google.com/translate_tts?ie=UTF-8&q=armadillo&tl=en&total=1&idx=0&textlen=9&tk=786739.706391&client=webapp ']
  ,['mouse or rat','https://translate.google.com/translate_tts?ie=UTF-8&q=mouse%20or%20rat&tl=en&to…l=1&idx=0&textlen=12&tk=4638.447610&client=webapp ']
  ,['Golden hamster','https://translate.google.com/translate_tts?ie=UTF-8&q=Golden%20hamster&tl=en&to…1&idx=0&textlen=14&tk=762087.878211&client=webapp ']
  ,['Whale','https://translate.google.com/translate_tts?ie=UTF-8&q=Whale&tl=en&total=1&idx=0&textlen=5&tk=896536.747644&client=webapp ']
  ,['Unicorn','https://translate.google.com/translate_tts?ie=UTF-8&q=Unicorn&tl=en&total=1&idx=0&textlen=7&tk=427230.17082&client=webapp ']
  ,['frog','https://translate.google.com/translate_tts?ie=UTF-8&q=frog&tl=en&total=1&idx=0&textlen=4&tk=66950.508898&client=webapp ']
  ,['Raccoon dog','https://translate.google.com/translate_tts?ie=UTF-8&q=Raccoon%20dog&tl=en&total=1&idx=0&textlen=11&tk=433122.23942&client=webapp ']
  ,['Chimpanzee','https://translate.google.com/translate_tts?ie=UTF-8&q=Chimpanzee&tl=en&total=1&idx=0&textlen=10&tk=954396.546424&client=webapp']
  ,['reindeer','https://translate.google.com/translate_tts?ie=UTF-8&q=reindeer&tl=en&total=1&idx=0&textlen=8&tk=331631.248075&client=webapp ']
  ,['ape','https://translate.google.com/translate_tts?ie=UTF-8&q=ape&tl=en&total=1&idx=0&textlen=3&tk=493703.85731&client=webapp ']
  ,['deer','https://translate.google.com/translate_tts?ie=UTF-8&q=deer&tl=en&total=1&idx=0&textlen=4&tk=453315.11431&client=webapp ']
  ,['Momonga','https://translate.google.com/translate_tts?ie=UTF-8&q=Momonga&tl=en&total=1&idx=0&textlen=7&tk=64163.408775&client=webapp ']
  ,['Dolphin','https://translate.google.com/translate_tts?ie=UTF-8&q=Dolphin&tl=en&total=1&idx=0&textlen=7&tk=802061.719721&client=webapp ']
  ,['Sheep','https://translate.google.com/translate_tts?ie=UTF-8&q=Sheep&tl=en&total=1&idx=0&textlen=5&tk=963294.554170&client=webapp ']
  ,['camel','https://translate.google.com/translate_tts?ie=UTF-8&q=camel&tl=en&total=1&idx=0&textlen=5&tk=179812.292864&client=webapp ']
  ,['Pig','https://translate.google.com/translate_tts?ie=UTF-8&q=Pig&tl=en&total=1&idx=0&textlen=3&tk=859288.776956&client=webapp ']
  ,['hyena','https://translate.google.com/translate_tts?ie=UTF-8&q=hyena&tl=en&total=1&idx=0&textlen=5&tk=28450.436550&client=webapp ']
  ,['tiger','https://translate.google.com/translate_tts?ie=UTF-8&q=tiger&tl=en&total=1&idx=0&textlen=5&tk=493959.85987&client=webapp ']
  ,['Polar bear','https://translate.google.com/translate_tts?ie=UTF-8&q=Polar%20bear&tl=en&total=1&idx=0&textlen=10&tk=152239.300235&client=webapp ']
  ,['Lion','https://translate.google.com/translate_tts?ie=UTF-8&q=Lion&tl=en&total=1&idx=0&textlen=4&tk=543756.952936&client=webapp ']
  ,['Tapir','https://translate.google.com/translate_tts?ie=UTF-8&q=Tapir&tl=en&total=1&idx=0&textlen=5&tk=518890.76942&client=webapp ']
  ,['Baboon','https://translate.google.com/translate_tts?ie=UTF-8&q=Baboon&tl=en&total=1&idx=0&textlen=6&tk=24291.432263&client=webapp ']
  ,['Wombat','https://translate.google.com/translate_tts?ie=UTF-8&q=Wombat&tl=en&total=1&idx=0&textlen=6&tk=584949.928401&client=webapp ']
  ,['seal','https://translate.google.com/translate_tts?ie=UTF-8&q=seal&tl=en&total=1&idx=0&textlen=4&tk=799909.717505&client=webapp ']
  ,['Bat','https://translate.google.com/translate_tts?ie=UTF-8&q=Bat&tl=en&total=1&idx=0&textlen=3&tk=770803.853143&client=webapp ']
  ,['Guinea pig','https://translate.google.com/translate_tts?ie=UTF-8&q=Guinea%20pig&tl=en&total=1&idx=0&textlen=10&tk=248747.331215&client=webapp']
  ,['Goat','https://translate.google.com/translate_tts?ie=UTF-8&q=Goat&tl=en&total=1&idx=0&textlen=4&tk=267596.185128&client=webapp ']
  ,['Sea otter','https://translate.google.com/translate_tts?ie=UTF-8&q=Sea%20otter&tl=en&total=1&idx=0&textlen=9&tk=669094.848834&client=webapp ']
  ,['Donkey','https://translate.google.com/translate_tts?ie=UTF-8&q=Donkey&tl=en&total=1&idx=0&textlen=6&tk=327057.146421&client=webapp ']
  ,['Amoeba','https://translate.google.com/translate_tts?ie=UTF-8&q=Amoeba&tl=en&total=1&idx=0&textlen=6&tk=899140.753184&client=webapp ']
  ,['Paramecium','https://translate.google.com/translate_tts?ie=UTF-8&q=Paramecium&tl=en&total=1&idx=0&textlen=10&tk=235432.348620&client=webapp ']
  ,['Euglena','https://translate.google.com/translate_tts?ie=UTF-8&q=Euglena&tl=en&total=1&idx=0&textlen=7&tk=24981.436209&client=webapp ']
  ,['Sea turtle','https://translate.google.com/translate_tts?ie=UTF-8&q=Sea%20turtle&tl=en&total=1&idx=0&textlen=10&tk=595007.1036891&client=webapp ']
  ,['snake','https://translate.google.com/translate_tts?ie=UTF-8&q=snake&tl=en&total=1&idx=0&textlen=5&tk=644176.987700&client=webapp ']
  ,['turtle','https://translate.google.com/translate_tts?ie=UTF-8&q=turtle&tl=en&total=1&idx=0&textlen=6&tk=686352.835444&client=webapp ']
  ,['crocodile','https://translate.google.com/translate_tts?ie=UTF-8&q=crocodile&tl=en&total=1&idx=0&textlen=9&tk=10833.453685&client=webapp ']
  ,['Hippo','https://translate.google.com/translate_tts?ie=UTF-8&q=Hippo&tl=en&total=1&idx=0&textlen=5&tk=349017.234813&client=webapp']
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


