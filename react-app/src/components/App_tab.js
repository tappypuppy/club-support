import React, { useState, useEffect } from 'react';
import Tab from 'react-bootstrap/Tab';
import Tabs from 'react-bootstrap/Tabs';
import AppCalendar from './Calendar';
import AddEvent from './Add_Event'

function ControlledTabsExample() {
    // ここわからん
    const [key, setKey] = useState('home');
    const [eventItems, setEventItems] = useState([{event_name:'aaa', event_time:'2022-08-26'}])
    console.log(eventItems[0])

    const addEventItem = (event_name,event_time) => {
        setEventItems([...eventItems, {event_name:event_name, event_time:event_time}])
    }
    

    // useEffect(() => {
    //     (async () => {
    //         setEventItems([
    //             {id:1,event_name:'あああ',event_time:'2022-08-26'},
    //         ])
    //     })()
    // }, [])

    return (
        <Tabs
        id="controlled-tab"
        activeKey={key}
        onSelect={(k) => setKey(k)}
        className="mb-3"
        >
        <Tab eventKey="home" title="Home">
            ホーム画面表示
        </Tab>
        <Tab eventKey="calender" title="Calender">
            カレンダー表示
            <AddEvent addEventItem={addEventItem}/>
            <AppCalendar eventItems={eventItems}/>
        </Tab>
        <Tab eventKey="menu" title="Menu">
            共有メニュー表示
        </Tab>
        </Tabs>
    );
}

export default ControlledTabsExample;