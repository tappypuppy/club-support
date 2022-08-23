import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid"; // pluginは、あとから
import interactionPlugin from "@fullcalendar/interaction";

function App_Calendar() {
  return (
      <FullCalendar 
      plugins={[dayGridPlugin]} 
      initialView="dayGridMonth" 
      locale="ja"
      />
  );
}

export default App_Calendar;