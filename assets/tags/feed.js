<feed>
  <button class="btn btn-primary" onclick={ this.doit }>Feed The Fur Ball!</button>

  <script>

  /*
    this.on('mount', function(){
        this.reload()
    })

    reload() {
        var obj = this;
        $.ajax({
            type: 'POST',
            url: '/feed',
            dataType: 'json',
        })
        .fail(function() {
            console.log("TODO: Handle failure better");
        })
        .done(function(data, textStatus, jqXhr) {
			console.log("done");
            obj.update()
        });
    }
	*/

    doit(e) {
        console.log(this.button_action)
        var obj = this
		console.log("feeding the fur ball");
        $.ajax({
            type: 'POST',
            url: 'feed'
        });
       obj.reload()
   }
  </script>
</feed>
