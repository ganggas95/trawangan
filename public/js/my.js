function readURL(input) {
	if (input.files && input.files[0]) {
		var reader = new FileReader();

		reader.onload = function (e) {
			$('#blah').attr('src', e.target.result);
		}
		reader.readAsDataURL(input.files[0]);
	}
    if (input.files && input.files[1]) {
        var reader = new FileReader();

        reader.onload = function (e) {
            $('#blah2').attr('src', e.target.result);
        }
        reader.readAsDataURL(input.files[1]);
    }
}
function readURL2(input) {
    if (input.files && input.files[0]) {
        var reader = new FileReader();

        reader.onload = function (e) {
            $('#blah2').attr('src', e.target.result);
        }
        reader.readAsDataURL(input.files[0]);
    }
}
$("#uploader").change(function(){
	readURL(this);
});
$("#uploader2").change(function(){
    readURL2(this);
});

$(document).on("click", ".delete-dialog", function () {
	var serviceId = $(this).data('id');
	$(".modal-body #idService").val( serviceId );
});
$('#editdialog').on("click", function () {
    var serviceId = $(this).data('id');
    var data = $(this).data('options');
    $(".modal-body #idService").val( serviceId );
});

$('#editDesc').on("click", function () {
    var serviceId = $(this).data('id');
    var descService = $(this).data('desc');
    $(".modal-body #idService").val( serviceId );
    $(".modal-body .form-group textarea#descText").val(descService);
});


