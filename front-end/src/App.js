import { useState, useEffect } from 'react'
import './App.css'

function App() {
  const [data, setData] = useState(null);

  useEffect(() => {
    (async() => {
      const serviceResponse = await fetch('http://localhost:8080/api/v1/pings/services');
      const serviceItems = await serviceResponse.json();

      const data = [];
      for (const serviceItem of serviceItems) {
        const pingResponse = await fetch(`http://localhost:8080/api/v1/pings/overview-hours?serviceName=${serviceItem}`);
        const pingItems = await pingResponse.json();

        data.push({
          serviceName: serviceItem,
          pings: pingItems
        });
      }
      setData(data);
    })();
  }, []);

  if (!data) {
    return (<div className="App">
      <header className="App-header">
        <span className="App-name">Monitier</span>
      </header>

      <div className="grid">
      <div>
        <div>Loading...</div>
        <div className="status">
        </div>

      </div>
      </div>
    </div>)
  }

  return (
    <div className="App">
      <header className="App-header">
        <span className="App-name">Monitier</span>
      </header>

      <div className="grid">
        {data.map((item, mainIndex) => {

          const statuses = item.pings.map((ping, index) => <div key={index.toString()} className={ping.averageResponseTime === 0 ? 'status-off' : 'status-day'}></div>)
          return <div key={mainIndex}><div>{item.serviceName}</div><div className="status">{statuses}</div></div>
        })}
      </div>
    </div>
  );
}

export default App;
