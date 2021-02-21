import React, { FC } from 'react';
import Textinfo from './textTreat'
import { Link as RouterLink } from 'react-router-dom';

import { Cookies } from 'react-cookie/cjs';
import { useEffect, useState } from 'react';
import { Avatar,TextField,Paper,makeStyles,Button} from '@material-ui/core';
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
import { EntTreatment } from '../../api/models/EntTreatment';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';

const cookies = new Cookies();
const Name = cookies.get('Name');
const Img = cookies.get('Img');

const findTreatment: FC<{}> = () => {
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
  
  const classes = useStyles();
  const [alert, setAlert] = React.useState(true);
  const [loading, setLoading] = React.useState(true);
  const [status, setStatus] = React.useState(false);

  const [Users, setUsers] = React.useState<Partial<EntUser>>();
  const [PatID, setPatID] = React.useState<number>(0);
  const [treatments, setTreatments] = useState<EntTreatment[]>(Array);
  const [Patientrecord, setPatientrecord] = useState<EntPatientrecord[]>([])

  const getPatientrecord = async () => {
    const res = await http.listPatientrecord({ offset: 0 });
    setPatientrecord(res);
  };

  const getTreatments = async () => {
    const res = await http.getTreatment({ id: PatID })
    setLoading(false);
    setTreatments(res);
    if (res.length != 0) {
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
    getPatientrecord();
    getImg();
  }, [loading]);

  
  const SearchTreatment = async () => {
    getTreatments();
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ส่วนบันทึกการรักษา`}>
        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
      <Content>
      <ContentHeader title="ค้นหาบันทึกการรักษา">
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
              SearchTreatment();
            }}
            variant="contained"
            color="primary"
            style={{backgroundColor: "#00af91",padding: '6px 12px',}}
          >
            ค้นหา
               </Button>&emsp;
            <Link component={RouterLink} to="/createTreatment">
            <Button variant="contained"
            color="secondary">
              บันทึกการรักษา
           </Button>           
          </Link>&emsp;
          <Link component={RouterLink} to="/">
            <Button variant="contained"
            color="default">
              กลับสู่หน้าหลัก
           </Button>           
          </Link>&emsp;
        </ContentHeader>
        <TableContainer component={Paper}>
                            <Table className={classes.table} aria-label="simple table">
                              <TableHead>
                                <TableRow>
                                  <TableCell align="center">เลขที่การรักษา</TableCell>
                                  <TableCell align="center">แพทย์</TableCell>
                                  <TableCell align="center">ผู้เข้ารับการรักษา</TableCell>
                                  <TableCell align="center">รูปแบบการรักษา</TableCell>
                                  <TableCell align="center">วันเวลาที่รักษา</TableCell>
                                  <TableCell align="center"></TableCell>
                                </TableRow>
                              </TableHead>

                              <TableBody>
                              {treatments.map(treatment => (
                              <TableRow>
                                <TableCell align="center">{treatment.id}</TableCell>
                                <TableCell align="center">{treatment.edges?.edgesOfDoctor?.edges?.edgesOfDoctorinfo?.doctorname} {treatment.edges?.edgesOfDoctor?.edges?.edgesOfDoctorinfo?.doctorsurname}</TableCell>
                                <TableCell align="center">{treatment.edges?.edgesOfPatientrecord?.name}</TableCell>
                                <TableCell align="center">{treatment.edges?.edgesOfTypetreatment?.typetreatment}</TableCell>
                                <TableCell align="center">{treatment.datetreat}</TableCell>
                                <TableCell align="center"><Textinfo id={treatment.id}></Textinfo></TableCell>  
                              </TableRow>
                            ))}
                              </TableBody>
                            </Table>
                          </TableContainer>
      </Content>
    </Page>
  );
};

export default findTreatment;
