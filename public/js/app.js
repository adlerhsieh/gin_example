$(document).ready(function(){
  $("#create-vacation").click(function(){
    var form = $("#vacation");
    var id = form.attr("data-id");
    $.post("/employees/"+id+"/vacations", form.serialize()).done(function(data){
      console.log(data);
    });
  });
});
