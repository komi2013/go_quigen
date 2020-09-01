
var drawerIsOpen = false;
$('#menu_td').click(function(){
  if(drawerIsOpen){
    $('#drawer').css({'left':'-100%'});
    drawerIsOpen = false;
  }else{
    $('#drawer').css({'left': '-1px','top':$(window).scrollTop()+51+'px'});
//    $('#ad_menu').empty().append(ad_menu_iframe);
    drawerIsOpen = true;
  }
});

if(localStorage.etoColor){
  $('#page_myimg').css({'background-color':'#'+localStorage.etoColor});
}
if(localStorage.myphoto){
  $('#page_myimg').attr('src',localStorage.myphoto); 
}
