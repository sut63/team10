import React, { useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import { DefaultApi } from'../../api/apis';
import { Typography, Avatar } from '@material-ui/core'
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save';
import UndoIcon from '@material-ui/icons/Undo';
import { Alert } from '@material-ui/lab';
import { Link as RouterLink } from 'react-router-dom';
import TextField from '@material-ui/core/TextField';

import { EntMedicalrecordstaff } from '../../api/models/EntMedicalrecordstaff';
import { EntGender } from '../../api/models/EntGender';
import { EntPrename } from '../../api/models/EntPrename';
import { EntUser } from '../../api/models/EntUser';

import { Cookies } from 'react-cookie/cjs';//cookie

// header css
const HeaderCustom = {
  minHeight: '50px',
};

const cookies = new Cookies();
const MEDID = cookies.get('Med'); 
const Name = cookies.get('Name');
const Img = cookies.get('Img');


const useStyles = makeStyles((theme: Theme) =>
    createStyles({
      button: {
        display: 'block',
        marginTop: theme.spacing(2),
      },
      formControl: {
          width: 200,
        },
        selectEmpty: {
          marginTop: theme.spacing(2),
        },
        paper: {
          marginTop: theme.spacing(2),
          marginBottom: theme.spacing(2),
        },
        textField: {
            marginLeft: theme.spacing(1),
            marginRight: theme.spacing(1),
            width: 200,
        },
    }),
  );

export  default  function Create() {
    
  const classes = useStyles();
  const api = new DefaultApi();

  const [status, setStatus] = React.useState(false);
  const [alert, setAlert] = React.useState(true);

  const [prename, setPrename] = React.useState<EntPrename[]>([]);
  const [gender, setGender] = React.useState<EntGender[]>([]);
  const [medicalrecordstaff, setMedicalrecordstaff] = React.useState<Partial<EntMedicalrecordstaff>>();
   
  const [prenameid, setprenameId] = React.useState(Number);
  const [genderid, setgenderId] = React.useState(Number);
  const [name, setname] = React.useState(String);
  const [idcardnumber, setidcardnumber] = React.useState(String);
  const [age, setage] = React.useState(String);
  const [bloodtype, setbloodtype] = React.useState(String);
  const [disease, setdisease] = React.useState(String);
  const [allergic, setallergic] = React.useState(String);
  const [phonenumber, setphonenumber] = React.useState(String);
  const [email, setemail] = React.useState(String);
  const [home, sethome] = React.useState(String);
  const [loading, setLoading] = React.useState(true);
  const [Users, setUsers] = React.useState<Partial<EntUser>>();

  useEffect(() => {
    const getPrename = async () => {
        const res = await api.listPrename({ limit: 5, offset: 2 });
        setLoading(false);
        setPrename(res);
    };
    getPrename();
    const getGender = async () => {
        const res = await api.listGender({ limit: 2, offset: 0 });
        setLoading(false);
        setGender(res);
    };
    getGender();
    const getMedicalrecordstaff = async () => {
        const res = await api.getMedicalrecordstaff({ id: Number(MEDID) });
        setLoading(false);
        setMedicalrecordstaff(res);
    };
    getMedicalrecordstaff();
    const getImg = async () => {
      const res = await api.getUser({ id: Number(Img) });
      setLoading(false);
      setUsers(res);
    };
    getImg();
    }, [loading]);
    
    //handle
    const PrenamehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setprenameId(event.target.value as number);
      };
      const GenderhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setgenderId(event.target.value as number);
      };
      const NamehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setname(event.target.value as string);
      };
      const IdcardnumberhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setidcardnumber(event.target.value as string);
      };
      const AgehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setage(event.target.value as string);
      };
      const BloodtypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setbloodtype(event.target.value as string);
      };
      const DiseasehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setdisease(event.target.value as string);
      };
      const AllergichandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setallergic(event.target.value as string);
      };
      const PhonenumberhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setphonenumber(event.target.value as string);
      };
      const EmailhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setemail(event.target.value as string);
      };
      const HomehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        sethome(event.target.value as string);
      };
    
      const CreatePatientrecord = async () => {
        if((name != '')&&(idcardnumber != '')&&(age != '')&&(bloodtype != '')&&(disease != '')&&(
          allergic != '')&&(phonenumber != '')&&(email != '')&&(home != '')&&(prenameid != null)&&(genderid != null)){
        const patientrecord = {
          prename : prenameid,
          gender : genderid ,
          medicalrecordstaff : medicalrecordstaff?.id,
          name : name,
          idcardnumber : idcardnumber,
          age : age,
          bloodtype : bloodtype,
          disease : disease,
          allergic : allergic,
          phonenumber : phonenumber,
          email : email,
          home : home,
        };
        console.log(patientrecord);
        const res: any = await api.createPatientrecord({ patientrecord : patientrecord });
        setStatus(true);
        if (res.id != '') {
          setAlert(true);
          setTimeout(() => {
            setStatus(false);
          }, 1000);
        }
       } 
       else {
        setStatus(true);
        setAlert(false);
        setTimeout(() => {
          setStatus(false);
        }, 1000);
        }
      };

    return(
        <Page theme={pageTheme.home}>
          <Header style={HeaderCustom} title={`ลงทะเบียนผู้ป่วยนอก`}>
            <Avatar alt="Remy Sharp" src={Users?.images as string} />
            <div style={{ marginLeft: 10 }}>{Name}</div>
          </Header>
            <Content>
            <br />
            <br />
            {status ? (
                <div>
                {alert ? (
                  <Alert severity="success">
                    บันทึกสำเร็จ !
                  </Alert>
                  ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
                  </Alert>
                )}
                </div>
              ) : null}
            <Typography variant="h6" gutterBottom align="center">
              <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>คำนำหน้าชื่อ</InputLabel>
              <Select
                name="prename"
                label="คำนำหน้าชื่อ"
                value={prenameid}
                onChange={PrenamehandleChange}
              >
              {prename.map(item => {
                return (
                  <MenuItem value={item.id}>
                  {item.prefix}
                  </MenuItem>
                  );
                })}
              </Select>
              </FormControl> &emsp;

            <TextField 
            label="ชื่อ-นามสกุล" 
            variant="outlined" 
            value={name}
            onChange={NamehandleChange}
            /> &emsp;

              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เพศ</InputLabel>
                  <Select
                    name="gender"
                    label="เพศ"
                    value={genderid}
                    onChange={GenderhandleChange}
                  >
                  {gender.map(item => {
                    return (
                      <MenuItem value={item.id}>
                      {item.genderstatus}
                      </MenuItem>
                      );
                    })}
                  </Select>
              </FormControl>

            <div className={classes.paper}></div>
            <TextField 
            label="เลขบัตรประจำตัวประชาชน" 
            variant="outlined" 
            type="number"
            value={idcardnumber}
            onChange={IdcardnumberhandleChange}
            /> &emsp;

            <TextField 
            label="อายุ" 
            variant="outlined" 
            type="number"
            value={age}
            onChange={AgehandleChange}
            /> &emsp;
            
            <TextField 
            label="กรุ๊ปเลือด" 
            name="bloodtype"
            variant="outlined" 
            value={bloodtype}
            onChange={BloodtypehandleChange}
            /> 
            
            <div className={classes.paper}></div>
            <TextField 
            name="disease"
            label="โรคประจำตัว" 
            variant="outlined" 
            style={{ width: "25ch"}}
            value={disease}
            onChange={DiseasehandleChange}
            /> &emsp;
            
            <TextField
            label="แพ้ยา"  
            name="allergic"
            variant="outlined" 
            style={{ width: "25ch"}}
            value={allergic}
            onChange={AllergichandleChange}
            /> 
            
            <div className={classes.paper}></div>
            <TextField 
            label="เบอร์โทรที่ติดต่อได้" 
            variant="outlined"
            style={{ width: "25ch"}} 
            value={phonenumber}
            onChange={PhonenumberhandleChange}
            /> &emsp;
            
            <TextField 
            label="อีเมล์" 
            variant="outlined"
            style={{ width: "25ch"}} 
            value={email}
            onChange={EmailhandleChange}
            /> 

            <div className={classes.paper}></div>
            <TextField
            label="ที่อยู่"
            multiline
            variant="outlined"
            style={{ width: "67ch"}}
            rows={3}
            value={home}
            onChange={HomehandleChange}
            />
              <br />
              <br /> พนักงานเวชระเบียน : {medicalrecordstaff?.name}
              <br />
              <br />
              <Typography variant="h6" gutterBottom  align="center">
              <Button
                onClick={() => {
                  CreatePatientrecord();
                }}
                variant="contained"
                color="primary"
                style={{backgroundColor: "#26c6da"}}
                startIcon={<SaveIcon />}
                >
                Submit
              </Button>
              <Button
                    style={{ marginLeft: 40 }}
                    startIcon={<UndoIcon />}
                    component={RouterLink}
                    to="/"
                    variant="contained"
                >
                  Back
                </Button>
                <Button
                    style={{ marginLeft: 40 }}
                    component={RouterLink}
                    to="/Patientrecord"
                    variant="contained"
                    color="secondary"
                >
                    SHOW
                </Button>
              </Typography>
            </Typography>
          </Content>
      </Page>
  );
}