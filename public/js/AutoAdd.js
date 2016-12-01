function addRow(section, initRow) {
    var newRow = initRow.clone().removeAttr('id').addClass('new').insertAfter(initRow);
    var deleteRow = $('<span class="input-group-addon"><button class="glyphicon glyphicon-remove" id="rowDelete" type="button" style="border:none;background:none;"></button></span>');
    newRow.find('#add').attr("id", "remove");
    $('#remove').remove();
    newRow.append(deleteRow).on('click', 'button#rowDelete', function() {
            removeRow(newRow);
        });
    newRow.slideDown("200", function() {
            $(this).find('.row')
                .find('div')
                .find('div')
                .find('input')
                .focus();
        });
}
                         
function removeRow(newRow) {
    newRow.slideUp("200", function() {
        $(this).next('div:not(#initRow)')
            .focus()
            .end()
            .end()
            .remove();
    });
}
    $(function () {
        var initRow = $('#initRow'),
            section = initRow.parent('section');
        $('#add').click(function(){
            addRow(section, initRow);
        })                           
    }); 


   


    