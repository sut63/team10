import React, { FC } from 'react';
import MoreInfo from './detail'
import { Link as RouterLink } from 'react-router-dom';

import { Cookies } from 'react-cookie/cjs';
import { useEffect, useState } from 'react';
import { Avatar, TextField, Paper, makeStyles, Button } from '@material-ui/core';
import { Autocomplete, Alert } from '@material-ui/lab';
import { Content, Header, Page, pageTheme, ContentHeader, Link, } from '@backstage/core';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';

import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
import { EntHistorytaking } from '../../api/models/EntHistorytaking';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';

const cookies = new Cookies();
const Name = cookies.get('Name');
const Img = cookies.get('Img');

const HistorytakingSearch: FC<{}> = () => {
  const http = new DefaultApi();
  const useStyles = makeStyles(theme => ({
    table: {
      minWidth: 650,
    },
    root: {
      '& > *': {
        borderBottom: 'unset',
      },
      width: '100%',
      maxWidth: 360,
      backgroundColor: theme.palette.background.paper,

    },
  }));

  const [PatID, setPatID] = React.useState<number>(0);
  const classes = useStyles();
  const [alert, setAlert] = React.useState(true);
  const [loading, setLoading] = React.useState(true);
  const [status, setStatus] = React.useState(false);
  const [Users, setUsers] = React.useState<Partial<EntUser>>();
  const [Historytakings, setHistorytakings] = useState<EntHistorytaking[]>(Array);
  const [Patientrecord, setPatientrecord] = React.useState<EntPatientrecord[]>([]);

  const getPatientrecord = async () => {
    const res = await http.listPatientrecord({ offset: 0 });
    setPatientrecord(res);
  };

  const Start = async () => {
    const res = await http.listHistorytaking({ offset: 0 });
    setHistorytakings(res);
  };

  const getHistorytakings = async () => {
    const res = await http.getHistorytaking({ id: PatID })
    setLoading(false);
    setHistorytakings(res);
    if (res.length != 0 || PatID === undefined || PatID === null) {
      setStatus(true);
      setAlert(true);
    } else {
      setStatus(true);
      setAlert(false);
    }
    setTimeout(() => {
      setStatus(false);
    }, 1000);

  };

  const handleChange = (event: any, value: unknown) => {
    console.log(value)
    Patientrecord.map(item => {
      if (item.name === value) {
        setPatID(item.id as number);
      } else if (value === null) {
        setPatID(0);
      }
    })
  };

  const TexthandleChang = (event: React.ChangeEvent<{ value: unknown; }>) => {
    console.log(event.target.value as string)
    Patientrecord.map(item => {
      if (item.name === event.target.value as string) {
        setPatID(item.id as number);
      } else if (event.target.value as string === null) {
        setPatID(0);
      }
    })
  }
  useEffect(() => {
    const getImg = async () => {
      const res = await http.getUser({ id: Number(Img) });
      setLoading(false);
      setUsers(res);
    };
    getImg();
    getPatientrecord();
    Start();

  }, [loading]);
  const SearchHistorytaking = async () => {
    getHistorytakings();
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`HISTORYTAKING DEPARTMENT`}>
        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
      <Content>
        <ContentHeader title="SHEACH HISTORYTAKING">
          <br />
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  พบข้อมูลผู้ป่วย
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    ไม่พบข้อมูลผู้ป่วย
                  </Alert>
                )}
            </div>
          ) : null}
        &emsp;
        <Autocomplete
            id="patientname"
            freeSolo
            options={Patientrecord.map(option => option.name)}
            onChange={handleChange}
            closeText='Close'
            renderInput={(params) => (
              <TextField {...params} label="ชื่อผู้ป่วย" margin="normal" variant="outlined" onChange={TexthandleChang} style={{ width: "100ch" }} />
            )}
          />
          &emsp;
          <Button
            onClick={() => {
              SearchHistorytaking();
            }}
            variant="contained"
            color="primary"
            style={{ backgroundColor: "#00af91", padding: '6px 12px', }}
          >
            SEARCH
         </Button>&emsp;
         <Link component={RouterLink} to="/">
            <Button variant="contained" color="primary" style={{ backgroundColor: "#00838f" }}>
              HOME
           </Button>
          </Link>&emsp;

         <Link component={RouterLink} to="/createHistorytaking">
            <Button variant="contained" color="primary" style={{ backgroundColor: "#006064" }}>
              ADD DATA
           </Button>
          </Link>
        </ContentHeader>
          <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">ID</TableCell>
                <TableCell align="center">Hight</TableCell>
                <TableCell align="center">Weight</TableCell>
                <TableCell align="center">T</TableCell>
                <TableCell align="center">P</TableCell>
                <TableCell align="center">R</TableCell>
                <TableCell align="center">BP</TableCell>
                <TableCell align="center">O2</TableCell>
                <TableCell align="center">Symptom</TableCell>
                <TableCell align="center">Nurse</TableCell>
                <TableCell align="center">Symptomseverity</TableCell>
                <TableCell align="center">Department</TableCell>
                <TableCell align="center">Patient</TableCell>
                <TableCell align="center">Datetime</TableCell>
                <TableCell align="center">Manage</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {Historytakings.map(item => (
                <TableRow>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.hight}</TableCell>
                  <TableCell align="center">{item.weight}</TableCell>
                  <TableCell align="center">{item.temp}</TableCell>
                  <TableCell align="center">{item.pulse}</TableCell>
                  <TableCell align="center">{item.respiration}</TableCell>
                  <TableCell align="center">{item.bp}</TableCell>
                  <TableCell align="center">{item.oxygen}%</TableCell>
                  <TableCell align="center">{item.symptom}</TableCell>
                  <TableCell align="center">{item.edges?.edgesOfNurse?.name}</TableCell>
                  <TableCell align="center">{item.edges?.edgesOfSymptomseverity?.symptomseverity}</TableCell>
                  <TableCell align="center">{item.edges?.edgesOfDepartment?.department}</TableCell>
                  <TableCell align="center">{item.edges?.edgesOfPatientrecord?.name}</TableCell>
                  <TableCell align="center">{item.datetime}</TableCell>
                  <TableCell align="center">
                    <MoreInfo id={item.id}></MoreInfo>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>  
      </Content>
    </Page>
  );
};

export default HistorytakingSearch;