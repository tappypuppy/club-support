import { useState } from "react"

function AddEvent(props){
    const [event_name, setEventName] = useState('')
    const [event_time, setEventTime] = useState('')

    const changeEventName = (e) =>{
        setEventName(e.target.value)
    }

    const changeEventTime = (e) =>{
        setEventTime(e.target.value)
    }

    const save = (e) => {
        e.preventDefault()
        console.log(event_name)
        console.log(event_time)
        props.addEventItem(event_name, event_time)
        setEventName('')
        setEventTime('')
        // setEvents({title: "event 3", start:"2022-08-26"})
    }
    return(
        <form>
            <input type="text" placeholder="event name" value={event_name} onChange={changeEventName}/>
            <input type="date" placeholder="event time" value={event_time} onChange={changeEventTime}/>
            <button onClick={save}>追加</button>
        </form>
    )
}

export default AddEvent