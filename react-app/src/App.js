import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import Container from 'react-bootstrap/Container';
import { Navbar, Nav} from 'react-bootstrap';
import ControlledTabsExample from './components/App_tab'
// import Sonnet from '../../components/Sonnet';


function App() {
  return (
    <div className="App">
      <header className="App-header">
        <Navbar bg="dark" variant="dark">
          <Container>
            <Navbar.Brand href="#home">TITLE</Navbar.Brand>
            <Nav>
              <Navbar.Collapse className='justify-content-end'>
                <Nav.Link href="#user">
                  UserName
                </Nav.Link>
              </Navbar.Collapse>
            </Nav>
            
          </Container>
        </Navbar>
      </header>
      <div>
        <ControlledTabsExample />
      </div>
      {/* tabs書く */}
    </div>
  );
}

export default App;
