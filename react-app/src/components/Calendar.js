import { useCallback } from "react";
import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid"; // pluginは、あとから
import interactionPlugin,  { DateClickArg } from "@fullcalendar/interaction";
import React, {useEffect, useState} from 'react'
import Modal from "./Modal"; //Modalコンポーネントをimportする





function AppCalendar(props) {
    const handleDateClick = useCallback((arg: DateClickArg) => {
        // alert(arg.dateStr);
        setShowModal(true);
        
        setday(arg.dateStr)
      }, []);
    const [day, setday] = useState();
    const [showModal, setShowModal] = useState(false);
    // const [events,setEvents] = useState([
    // { title: "event 1", start:"2022-08-24"},
    // { title: "event 2", start:"2022-08-25"},
    // { title: props.eventItems[0].event_name, start: props.eventItems[0].event_time}
    // ])
    let event_array = [
      { title: "event 1", start:"2022-08-24"},
      { title: "event 2", start:"2022-08-25"},];
    for (let step = 0; step < props.eventItems.length; step++) {
      event_array.push({ title: props.eventItems[step].event_name, start: props.eventItems[step].event_time});
    }
    const [events,setEvents] = useState(event_array)
    console.log("events:");
    console.log(events);

    console.log(props)

    const ShowModal = () => {
      setShowModal(true);
    };

  return (
    <>
    <FullCalendar 
    plugins={[dayGridPlugin,interactionPlugin]} 
    initialView="dayGridMonth" 
    locale="ja" 
    events={events}
    dateClick={handleDateClick}

    
    />
    <button onClick={ShowModal}>Open Modal</button>

    <Modal showFlag={showModal} setShowModal={setShowModal} content={day} />


    </>
  );


  
}

export default AppCalendar;