import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import { Content, Header, Page, pageTheme, ContentHeader, } from '@backstage/core';

import {

  FormControl,
  Select,
  InputLabel,
  MenuItem,
  Avatar,

  Link,
  Button,
} from '@material-ui/core';
import Timer from '../Timer';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { ControllersPatientrights } from '../../api/models/ControllersPatientrights';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';
import { EntPatientrightstype } from '../../api/models/EntPatientrightstype';
import { EntInsurance } from '../../api/models/EntInsurance';
import { EntMedicalrecordstaff } from '../../api/models/EntMedicalrecordstaff';
import { Alert } from '@material-ui/lab';
import { Cookies } from 'react-cookie/cjs';//cookie
import { EntUser } from '../../api/models/EntUser';


import { Link as RouterLink } from 'react-router-dom';

import { Image3Base64Function } from '../../image/Image3';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

const cookies = new Cookies();
const Name = cookies.get('Name');

// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
    display: 'flex',
    justifyContent: 'center',
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    margin: theme.spacing(3),
    width: 320,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: '25ch',
  },


  margin: {
    margin: theme.spacing(3),
  },
  input: {
    display: 'none',
  },

}));


const NewPatientright: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'สิทธ์' };
  const http = new DefaultApi();
  const cookies = new Cookies();

  const [Patientrights, setPatientrights] = React.useState<Partial<ControllersPatientrights>>({});

  const [Patientrightstype, setPatientrightstype] = React.useState<EntPatientrightstype[]>([]);
  const [Patientrecord, setPatientrecord] = React.useState<EntPatientrecord[]>([]);

  const [Insurance, setInsurance] = React.useState<EntInsurance[]>([]);
  const [Medicalrecordstaff, setMedicalrecordstaff] = React.useState<Partial<EntMedicalrecordstaff>>();
  const [status, setStatus] = React.useState(false);
  const [alert, setAlert] = React.useState(true);

  const UserId = cookies.get('ID');
  const med = cookies.get('Med');


  const getChangeOfUser = async () => {

    const name = "medicalrecordstaff";
    const value = parseInt(med, 10);
    setPatientrights({ ...Patientrights, [name]: value });
  };



  const getMedicalrecordstaffs = async () => {

    const res = await http.getMedicalrecordstaff({ id: Number(med) });
    setMedicalrecordstaff(res);
  };




  const getPatientrightstype = async () => {
    const res = await http.listPatientrightstype({ limit: 10, offset: 0 });
    setPatientrightstype(res);
  };

  const getPatientrecord = async () => {
    const res = await http.listPatientrecord({ limit: 10, offset: 0 });
    setPatientrecord(res);
  };

  const getInsurance = async () => {
    const res = await http.listInsurance({ limit: 10, offset: 0 });
    setInsurance(res);
  };




  // Lifecycle Hooks
  useEffect(() => {
    getChangeOfUser();
    getMedicalrecordstaffs();
    getPatientrightstype();
    getPatientrecord();
    getInsurance();
    getChangeOfUser();



  }, []);

  // set data to object Patientright
  const handleChange = (

    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof NewPatientright;
    const { value } = event.target;
    setPatientrights({ ...Patientrights, [name]: value });
  };





  const CreatePatientright = async () => {

    if ((Patientrights.insurance != null) && (Patientrights.medicalrecordstaff != null)
      && (Patientrights.patientrecord != null) && (Patientrights.patientrightstype != null)) {

      const res: any = await http.createPatientrights({
        patientrights: Patientrights


      });
      console.log(Patientrights);



      if (res.id != '') {
        setStatus(true);
        setAlert(true);
        setTimeout(() => {
          setStatus(false);
        }, 2000);
      }

    }
    else {
      setStatus(true);
      setAlert(false);
      setTimeout(() => {
        setStatus(false);
      }, 2000);
    }
  };



  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ลงทะเบียนสิทธิ์`}>
        <Avatar alt="Remy Sharp" src={Image3Base64Function} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
      <Content>
        <ContentHeader title="ข้อมูล">

          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  บันทึกสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    บันทึกไม่สำเร็จ
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>


        <div className={classes.root}>


          <form noValidate autoComplete="off">



            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>ผู้ป่วย</InputLabel>
              <Select
                name="patientrecord"
                value={Patientrights.patientrecord}
                onChange={handleChange}
              >
                {Patientrecord.map((item: any) => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.name}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>


          </form>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">

            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>ประกัน</InputLabel>
              <Select
                name="insurance"
                value={Patientrights.insurance}
                onChange={handleChange}
              >
                {Insurance.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.insurancecompany}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>



          </form>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">

            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>รูปแบบสิทธิ์</InputLabel>
              <Select
                name="patientrightstype"
                value={Patientrights.patientrightstype}
                onChange={handleChange}
              >
                {Patientrightstype.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.permission}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>


          </form>
        </div>



        <div className={classes.root}>
          <form noValidate autoComplete="off">

            <FormControl variant="outlined" className={classes.formControl}>

              <TextField
                disabled
                id="outlined-disabled"
                label="พนักงานเวชระเบียง"
                defaultValue=" "
                value={Medicalrecordstaff?.name as string}

                variant="outlined"
              />

            </FormControl>







          </form>

        </div>











        <div className={classes.root}>
          <form noValidate autoComplete="off">


            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreatePatientright();
                }}
                variant="contained"
                color="primary"
              >
                Submit
             </Button>
              <Button
                style={{ marginLeft: 40 }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                Back
                </Button>
              <Button
                style={{ marginLeft: 40 }}
                component={RouterLink}
                to="/Table_patientrights"
                variant="contained"
                color="secondary"
              >
                SHOW
                </Button>

            </div>
          </form>
        </div>

      </Content>
    </Page>
  );
};
export default NewPatientright;