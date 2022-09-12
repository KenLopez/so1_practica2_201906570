import React from 'react'
import { Col, Container, Row } from 'react-bootstrap'

function Counts(props) {
  return (
    <Container className='card mt-5 py-3 px-5'>
        <h3 className='mb-0'>Datos de Procesos</h3>
        <hr/>
        <Row>
            <Col>
            <h5>Procesos en ejecución: {props.ejecucion}</h5>
            <h5>Procesos detenidos: {props.detenidos}</h5>
            <h5>Procesos suspendidos: {props.suspendidos}</h5>
            <h5>Procesos zombie: {props.zombie}</h5>
            <h5><strong>Procesos totales: {props.total}</strong></h5>
            </Col>
        </Row>
    </Container>
  )
}

export default Counts