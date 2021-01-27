import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, Select, InputLabel, MenuItem, TextField, Button, InputAdornment, Grid } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
//api
import { DefaultApi } from '../../api/apis';
//table
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
//entity
import { EntTreatment } from '../../api/models/EntTreatment';
import { EntUser } from '../../api/models/EntUser';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';
import { EntDoctor } from '../../api/models/EntDoctor';
//alert
import Swal from 'sweetalert2'
//icon
import CancelTwoToneIcon from '@material-ui/icons/CancelTwoTone';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import DeleteTwoToneIcon from '@material-ui/icons/DeleteTwoTone';
import { Avatar } from '@material-ui/core';
import { Cookies } from 'react-cookie/cjs';//cookie


const cookies = new Cookies();
const Name = cookies.get('Name');
const Doc = cookies.get('Doc');
const Img = cookies.get('Img');


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    button: {
      display: 'block',
      marginTop: theme.spacing(2),
    },
    formControl: {
      width: 400,
    },
    selectEmpty: {
      marginTop: theme.spacing(1),
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
    table: {
      minWidth: 650,
    },
  }),
);

// header css
const HeaderCustom = {
  minHeight: '50px',
};

const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});


export default function ComponentsTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);
  
  const [Users, setUsers] = React.useState<Partial<EntUser>>();

  const [checkname, setcheckname] = useState(false);
  const [treatment, setTreatment] = useState<EntTreatment[]>([])
  const [doctor, setDoctor] = useState<EntDoctor[]>([])
  const [patientrecord, setPatientrecord] = useState<EntPatientrecord[]>([])


  const [name, setname] = useState(String);
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getTreatment = async () => {
      const res = await http.listTreatment({ offset: 0 });
      setLoading(false);
      setTreatment(res);
    };
    const getDocdor = async () => {
      const res = await http.listDoctor({ offset: 0 });
      setLoading(false);
      setDoctor(res);
    };
    const getPatientrecord = async () => {
      const res = await http.listPatientrecord({ limit: 2, offset: 0 });
      setLoading(false);
      setPatientrecord(res);
  };
    getTreatment();
    getDocdor();
    getPatientrecord();
  }, [loading]);

  const patientrecordhandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setcheckname(false);
    setname(event.target.value as string);

  };

  const cleardata = () => {
    setname("");
    setSearch(false);
    setcheckname(false);
    setSearch(false);

  }

  const checkresearch = async () => {
    var check = false;
    treatment.map(item => {
      if (name != "") {
        if (item.edges?.edgesOfPatientrecord?.name?.includes(name)) {
          setcheckname(true);
          alertMessage("success", "ค้นหาข้อมูลบุคลากรสำเร็จ");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ค้นหาข้อมูลบุคลากรไม่สำเร็จ");
    }
    console.log(checkname)
    if (name == "") {
      alertMessage("info", "แสดงข้อมูลบุคลากรทั้งหมดทั้งหมดในระบบ");
    }
  };

  return (

    <Page theme={pageTheme.home}>
        <Header style={HeaderCustom} title={`Treatment Department`}>
          <Avatar alt="Remy Sharp" src={Users?.images as string} />
          <div style={{ marginLeft: 10 }}>{Name}</div>
        </Header>
      <Content>
        <ContentHeader title="ค้นหาบันทึกการรักษา">
          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button
            onClick={() => {
              checkresearch();
              setSearch(true);
            }}
            variant="contained"
            color="secondary"
            startIcon={<SearchTwoToneIcon />}
          >
            ค้นหาข้อมูล
          </Button>
          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button
            onClick={() => {
              cleardata();
            }}
            variant="contained"
            startIcon={<DeleteTwoToneIcon />}
          >
            เคลียร์ข้อมูล
          </Button>
          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button
            href="/Personalwelcome"
            variant="contained"
            color="primary"
            startIcon={<CancelTwoToneIcon />}
          >
            ย้อนกลับ
          </Button>
        </ContentHeader>

        <div >
          <form noValidate autoComplete="off">
            <FormControl
              fullWidth
              variant="outlined"
              size="small"
            >
              <div ><strong>กรอก "ชื่อ" เพื่อทำการค้นหา</strong></div>
              <TextField
                id="name"
                variant="outlined"
                color="primary"
                type="string"
                size="small"
                value={name}
                onChange={patientrecordhandlehange}
              />
            </FormControl>
          </form>
        </div>

        <Grid container justify="center">
          <Grid item xs={12} md={10}>
            <Paper>
              {search ? (
                <div>
                  {  checkname ? (
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                            <TableCell align="center">เลขที่การรักษา</TableCell>
                            <TableCell align="center">แพทย์</TableCell>
                            <TableCell align="center">ผู้เข้ารับการรักษา</TableCell>
                            <TableCell align="center">รูปแบบการรักษา</TableCell>
                            <TableCell align="center">วันเวลาที่รักษา</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>

                        {treatment.filter((filter: any) => filter.edges?.edgesOfPatientrecord?.name.includes(name)).map((item: any) => (doctor.filter(t => t.id === item.edges?.edgesOfDoctor?.id).map((item2: any) =>(
                            <TableRow key={item.id}>
                              <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item2.edges?.edgesOfDoctorinfo?.doctorname} {item2.edges?.edgesOfDoctorinfo?.doctorsurname}</TableCell>
                              <TableCell align="center">{item.edges?.edgesOfPatientrecord?.name}</TableCell>
                              <TableCell align="center">{item.edges?.edgesOfTypetreatment?.typetreatment}</TableCell>
                              <TableCell align="center">{item.datetreat}</TableCell>
                            </TableRow>
                          ))))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : name == "" ? (
                      <div>
                        <TableContainer component={Paper}>
                          <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                              <TableRow>
                                <TableCell align="center">เลขที่การรักษา</TableCell>
                                <TableCell align="center">แพทย์</TableCell>
                                <TableCell align="center">ผู้เข้ารับการรักษา</TableCell>
                                <TableCell align="center">รูปแบบการรักษา</TableCell>
                                <TableCell align="center">วันเวลาที่รักษา</TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>

                            
                            {treatment.map((item: any) => (doctor.filter(t => t.id === item.edges?.edgesOfDoctor?.id).map((item2: any) => (
                                <TableRow key={item.id}>
                                  <TableCell align="center">{item.id}</TableCell>
                                  <TableCell align="center">{item2.edges?.edgesOfDoctorinfo?.doctorname} {item2.edges?.edgesOfDoctorinfo?.doctorsurname}</TableCell>
                                  <TableCell align="center">{item.edges?.edgesOfPatientrecord?.name}</TableCell>
                                  <TableCell align="center">{item.edges?.edgesOfTypetreatment?.typetreatment}</TableCell>
                                  <TableCell align="center">{item.datetreat}</TableCell>
                                </TableRow>
                            ))))}
                            </TableBody>
                          </Table>
                        </TableContainer>

                      </div>
                    ) : null}
                </div>
              ) : null}
            </Paper>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
}