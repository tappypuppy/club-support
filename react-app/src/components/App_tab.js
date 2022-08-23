import React, { useState } from 'react';
import Tab from 'react-bootstrap/Tab';
import Tabs from 'react-bootstrap/Tabs';
import App_Calendar from './Calendar';

function ControlledTabsExample() {
    // ここわからん
    const [key, setKey] = useState('home');

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
            <App_Calendar />
        </Tab>
        <Tab eventKey="menu" title="Menu">
            共有メニュー表示
        </Tab>
        </Tabs>
    );
}

export default ControlledTabsExample;