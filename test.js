var events = require('events');
var emiter = new events.EventEmitter;
function myEE() {
  events.EventEmitter.call(this);
  this.emit('run', 'This is nodejs');
}
emiter.on("run", function(str) {
  console.log(str)
})

var ins = new myEE();

