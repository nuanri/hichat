var React = require('react');
var PropTypes = React.PropTypes;

var NiceDate = React.createClass({
  render: function() {
    var date = new Date(this.props.date);
    var Y = date.getFullYear()
    var M = date.getMonth() + 1
    var D = date.getDate()
    var h = date.getHours()
    var m = date.getMinutes()
    var s = date.getSeconds()

    if (M < 10) {
      M = "0" + M
    }

    if (D < 10) {
      D = "0" + D
    }

    if (h < 10) {
      h = "0" + h
    }

    if (m < 10) {
      m = "0" + m
    }

    let dt_t = h+":"+m
    // var dt_d = M+"月"+D+"日"
    let dt_d = `${Y}-${M}-${D}`

    return (
      <span className="dt">{dt_d}<span className="t">{dt_t}</span></span>
    )
  }
})

export { NiceDate }
