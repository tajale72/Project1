import React, {Component} from 'react'


class Welcome extends Component {
    constructor() {
        super()
        this.state = {
            message: "welcome visitor",
        }
        this.state = {
            count : 1
        }
    }

    changeMessage() {
        this.setState({
            message : 'thanks for subsrcibing',
        })
    }
    changeMessage1() {
        this.setState({
            count :  2
        })
    }
    render() {
        return(
            <div>
            <h1>
               {this.state.message}
               {this.state.count}
               <button onClick ={this.changeMessage()} >Subscribe</button>
               <button onClick ={this.changeMessage1()} >count</button>
            </h1>

            </div>
            
        )
    }
}

export default Welcome