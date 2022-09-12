import { useEffect, useState } from 'react';
import {Col, Container, Navbar, Row} from 'react-bootstrap'
import axios from "axios";
import Tree from './components/Tree';

function App() {
  const [ejecucion, setEjecucion] = useState(0);
  const [detenidos, setDetenidos] = useState(0);
  const [suspendidos, setSuspendidos] = useState(0);
  const [zombie, setZombie] = useState(0);
  const [total, setTotal] = useState(0);
  const [request, setRequest] = useState(false);
  const [procs, setProcs] = useState([]);
  const [ram, setRam] = useState(0);
  const [cpu, setCpu] = useState(0);

  useEffect( () => {
    const getData = async () => {
      let res = await Promise.all([
        axios.get('http://34.125.72.118:3000/data/process', {
          headers: {
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          }
        }),
        axios.get('http://34.125.72.118:3000/data', {
          headers: {
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          }
        }),
      ])

      const data = res[0].data;
      setEjecucion(data.ejecucion);
      setSuspendidos(data.suspendidos);
      setDetenidos(data.detenidos);
      setZombie(data.zombie);
      setTotal(data.total);
      setProcs(data.procs);

      const info = res[1].data;
      setCpu(info.cpu);
      setRam(info.ram)

      const timer = setTimeout(() => {
        setRequest(!request)
      }, 3000);
      return () => clearTimeout(timer);
    }

    getData();

  },[request])

  return (
    <div className='App'>
      <Navbar bg='dark' variant='dark'>
        <Container fluid>
            <Navbar.Brand>
                Kenneth Haroldo López López - 201906570
            </Navbar.Brand>
        </Container>
      </Navbar>
      <Container className='card mt-5'>
        <Tree rows={procs}/>
      </Container>
      <Container className='card mt-5 py-3 px-5'>
        <h3 className='mb-0'>Datos de Procesos</h3>
        <hr/>
        <Row>
          <Col>
            <h5>Procesos en ejecución: {ejecucion}</h5>
            <h5>Procesos detenidos: {detenidos}</h5>
            <h5>Procesos suspendidos: {suspendidos}</h5>
            <h5>Procesos zombie: {zombie}</h5>
            <h5><strong>Procesos totales: {total}</strong></h5>
          </Col>
        </Row>
      </Container>
    </div>
  );
}

export default App;
