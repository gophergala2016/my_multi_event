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
      $.ajax({
        method: "POST",
        url: "/create_response?event_id="+event_id,
        data: JSON.stringify({answers: data}),
        success: function(resp) {
          $("#done").show();
          $("#questions").hide();
        },
        dataType: "json"
      });
      $("#questions_next").disable();
    }
  });
  if ($("#instructions").length == 0) {
    $("#questions").show();
  }
});
