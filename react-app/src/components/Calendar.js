import { useCallback } from "react";
import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid"; // pluginは、あとから
import interactionPlugin,  { DateClickArg } from "@fullcalendar/interaction";
import React, {useState,useEffect} from 'react'
import Modal from "./Modal"; //Modalコンポーネントをimportする

import axios from "axios";  // GinのREST APIを叩くためのモジュール

function AppCalendar(props) {
//-------------------------ReactでGinのAPIを叩く方法例-------------
  // 必要なURLを設定
  const URL = "http://localhost:8080/detail_user/1";
  // 取得した値を保存するusestate定義
  const [user, setUser] = useState(null);
  // リクエスト
  useEffect(() => {
    axios.get(URL).then((response) => {
      setUser(response.data.user.Name);
      console.log("respnse")
      console.log(response.data.user.Name)
    });
    }, []);

  console.log("user:")
  const user_name = String(user);
  console.log(user)
// ----------------------------------------------------------------------


    const handleDateClick = useCallback((arg: DateClickArg) => {
        // alert(arg.dateStr);
        setShowModal(true);
        
        setday(arg.dateStr)
      }, []);
    const [day, setday] = useState();
    const [showModal, setShowModal] = useState(false);
  
    let event_array = [
      { title: "event 1", start:"2022-08-24"},
      { title: "event 2", start:"2022-08-25"},];
    for (let step = 0; step < props.eventItems.length; step++) {
      event_array.push({ title: props.eventItems[step].event_name, start: props.eventItems[step].event_time});
    }
    // stateの更新はレンダー後らしい？
    // const [events,setEvents] = useState(event_array)
    // console.log("events:");
    // console.log(events);

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
    events={event_array}
    dateClick={handleDateClick}

    
    />

    <h2>{user_name}さん</h2>
    <button onClick={ShowModal}>Open Modal</button>

    <Modal showFlag={showModal} setShowModal={setShowModal} content={day} />

    
  

    </>
  );


  
}

export default AppCalendar;