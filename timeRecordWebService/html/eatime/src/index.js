import React from 'react'
import ReactDOM from 'react-dom'

const getLastTimeApi = '/api/lastEatTime';
const updateEatTimeApi = '/api/eat';

function EatButton() {
    return (
        <div className="eat button container" style={{textAlign: 'center'}}>
            <button id="eat-submit" className="ui button" onClick={eatProcedure}> 
                Eat
            </button>
        </div>
    );
}

function nowUnixTimestamp(){
    return Date.now()/1000;
}

function eatProcedure(){
    submitEat();
}


function submitEat(){
    fetch(updateEatTimeApi, {mode:'no-cors'}).then((res)=>{
        setTimeout(() => {
            window.location.reload();
        }, 500);
    }).catch((err) => {
        alert(err);
    });
}

function durationToString(sec_num) {
    let days = Math.floor(sec_num / 86400);
    let hours   = Math.floor(sec_num / 3600);
    let minutes = Math.floor((sec_num - (hours * 3600)) / 60);
    let seconds = sec_num - (hours * 3600) - (minutes * 60);
    return days + ' day ' + hours + ' hour ' + minutes + ' minute ' + parseInt(seconds) + ' second';
}

class App extends React.Component {
    constructor(props){
        super(props);

        this.state = {timeDiffString: '0 day 0 hour 0 second', lastEatTimestamp: null};
    }

    componentDidMount(){
        this.updateLastEatTimeDiffSeconds();
    }

    componentDidUpdate(){
        setTimeout(() => {
            this.updateTimeDiffString();
        }, 100);
    }

    updateTimeDiffString(){
        const durationSecond = nowUnixTimestamp() - this.state.lastEatTimestamp;
        this.setState({timeDiffString: durationToString(durationSecond)})
    }

    updateLastEatTimeDiffSeconds(){
        fetch(getLastTimeApi, {mode:'no-cors'}).then((res) => {
            return res.json();
        }).then((data) => {
            this.setState({lastEatTimestamp: parseInt(data['timeStamp'])});
            //console.log(this.state)
        }).catch((err)=>{
            console.log(err);
        });
    }

    render() {
        return (
            <div className="ui card">
                <div className="content">
                    <div className="header">
                        <div className="time diff container">{this.state.timeDiffString}</div>
                    </div>
                    <div style={{height: "500px", display: "block"}}></div>
                    <EatButton />
                </div>
            </div>
        );
    }
}

ReactDOM.render(<App />, document.querySelector("#root"));