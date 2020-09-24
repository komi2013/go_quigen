

$('input[name=choices]:eq(0)').prop('checked', true);

if($("#table-01 > tbody > tr:nth-child(4) > td > p:nth-child(2) > b").text()){
  console.log( $("#table-01 > tbody > tr:nth-child(4) > td > p:nth-child(2) > b").text() );
  console.log( $('input[name=choices]:eq(0)').val() );
  console.log( $('input[name=choices]:eq(1)').val() );
  console.log( $('input[name=choices]:eq(2)').val() );
  console.log( $('input[name=choices]:eq(3)').val() );
  console.log( $("#table-01 > tbody > tr:nth-child(2) > td > pre").text() );
  console.log( $("#table-01 > tbody > tr:nth-child(4) > td > pre").text() );  

  var q = localStorage.q ? JSON.parse(localStorage.q) : [];
  var quiz = [];
  var cho = [];
  cho.push($("#table-01 > tbody > tr:nth-child(4) > td > p:nth-child(2) > b").text());
  if($("#table-01 > tbody > tr:nth-child(4) > td > p:nth-child(2) > b").text() != $('input[name=choices]:eq(0)').val()){
    cho.push($('input[name=choices]:eq(0)').val());
  }
  if($("#table-01 > tbody > tr:nth-child(4) > td > p:nth-child(2) > b").text() != $('input[name=choices]:eq(1)').val()){
    cho.push($('input[name=choices]:eq(1)').val());
  }
  if($("#table-01 > tbody > tr:nth-child(4) > td > p:nth-child(2) > b").text() != $('input[name=choices]:eq(2)').val()){
    cho.push($('input[name=choices]:eq(2)').val());
  }
  if($("#table-01 > tbody > tr:nth-child(4) > td > p:nth-child(2) > b").text() != $('input[name=choices]:eq(3)').val()){
    cho.push($('input[name=choices]:eq(3)').val());
  }
  for (i = 0; i < cho.length; i++) {
    quiz[i] = cho[i];
  }
  quiz[4] = $("#table-01 > tbody > tr:nth-child(2) > td > pre").text();
  quiz[5] = $("#table-01 > tbody > tr:nth-child(4) > td > pre").text();
  q[q.length] = quiz;
  localStorage.q = JSON.stringify(q);
}

if(!$("#table-01 > tbody > tr:nth-child(4) > td > p:nth-child(2) > b").text()){
  setTimeout(function(){
    $('.button02')[0].click();
  },200);
}else{
  if(q.length < 210){
    setTimeout(function(){
      $('input[name=new_question]').click();
    },200);
  }
}
