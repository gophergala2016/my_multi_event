$(document).ready(function() {
  $("#instruction_next").click(function(){
    $("#instructions").hide();
    $("#questions").show();
  });
  $("#questions_next").click(function() {
    $("#questions").hide();
    $("#review").show();
  });
  $("#review_do").click(function() {
    $("#review_do").disable();
    debugger
    $.post({})
  })
  if ($("#instructions").length == 0) {
    $("#questions").show();
  }
});
