$(document).ready(function() {
  $("#instruction_next").click(function(){
    $("#instructions").hide();
    $("#questions").show();
  });
  $("#questions_next").click(function() {
    var valid = true;
    $(".validate").each(function(_, el) {
      if (el.value == "") {
        el.className += " invalid";
        valid = false;
      }
    })
    if (valid) {
      $("#questions").hide();
      var event_id = location.search.split('=')[1];
      var data = [];
      $("#schedule").serializeArray().forEach(function(el){ data.push(el.value) })
      $.post(
        "/create_response?event_id="+event_id,
        data,
        function(resp) {
          $("#done").show();
          $("#questions").hide();
        },
        "json"
      );
      $("#questions_next").disable();
    }
  });
  if ($("#instructions").length == 0) {
    $("#questions").show();
  }
});
