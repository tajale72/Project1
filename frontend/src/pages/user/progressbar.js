import React from 'react'



const Progress_bar = ({bgcolor,progress,height}) => {
    // const Status = {
	// 	height: '100px',
	// 	width: '100%',
	// 	backgroundColor: 'White',
    //     borderRadius: 50,
	// }
	
	const Parentdiv = {
		height: '30px',
		width: '100%',
		backgroundColor: 'whitesmoke',
        alignItems: 'center',

	}
	
	const Childdiv = {
		height: '100%',
		width: `${progress}%`,
		backgroundColor: bgcolor,
        borderRadius:40,
		textAlign: 'right'
	}
	
	const progresstext = {
		padding: 10,
		color: 'black',
		fontWeight: 900
	}
		
	return (
     <>
       <div style={Parentdiv}>
        
        <div style={Childdiv}>
            <span style={progresstext}>{`${progress}%`}</span>
        </div>
        </div>
     </> 

	)
}

export default Progress_bar;
