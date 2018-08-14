var resumeID = 1; // default, used to load main resume
new Vue({
    el: '#app',
    data:{
        info: {
            person_info: {
                name: "Loading..."
            }
        },
        educations: [],
        employments: [],
        volunteers:[],
        isEditing: false,
        isLoggedIn: false,
        input: {
            username: '',
            password: ''
        },
        token: '',
        invalidCred: false
    },
    methods: {
        loginButton: function() {
            if(this.input.username !== "" && this.input.password !== "") {
                self = this;
                axios
                    .post('/api/login', this.input)
                    .then(function(response) {
                        this.token = response.data;
                        if(this.token != null){
                            let modal = document.getElementById('id01');
                            let body = document.body;

                            modal.style.display = "none";
                            body.classList.remove("show-overlay");

                            self.isLoggedIn = true;
                        }
                    })
                    .catch(function (error) {
                    //console.log(error.response.status);
                        if(error.response.status === 401) {
                            self.invalidCred = true;
                        }
                    });
            }
            else {
                this.invalidCred = true;
            }
        },
        editButton: function() {
            this.isEditing = true;
        },
        saveButton: function() {
            axios
                .post('/api/resumejson/' + resumeID, this.info) //the data to post
                .then(r => console.log('r: ', JSON.stringify(r, null, 2)));
            this.isEditing = false;
        },
        addEduNote: function (idx) {
            console.log(this.info.educations);
            this.info.educations[idx].notes.push('');
        },
        deleteEduNote: function (i, j) {
            console.log(j);
            console.log(this.info.educations);
            this.info.educations[i].notes.splice(j, 1);
        },
        addJobNote: function (idx) {
            console.log(this.info.employments);
            this.info.employments[idx].notes.push('');
        },
        deleteJobNote: function (i, j) {
            console.log(j);
            console.log(this.info.employments);
            this.info.employments[i].notes.splice(j, 1);
        },
        addVolNote: function (idx) {
            console.log(this.info.volunteers);
            this.info.volunteers[idx].notes.push('');
        },
        deleteVolNote: function (i, j) {
            console.log(j);
            console.log(this.info.volunteers);
            this.info.volunteers[i].notes.splice(j, 1);
        },
        addEducation: function()  {
            var newEdu = {
                resume_id: resumeID,
                education_key: "",
                school: "",
                date_attended: "",
                is_deleted: false,
                notes: []
            };
            this.info.educations.push(newEdu);
        },
        addJob: function()  {
            var newJob = {
                resume_id: resumeID,
                employment_key: "",
                company: "",
                position: "",
                date_attended: "",
                is_deleted: false,
                notes: []
            };
            this.info.employments.push(newJob);
        },
        addVol: function()  {
            var newVol = {
                resume_id: resumeID,
                volunteer_key: "",
                company: "",
                position: "",
                date_attended: "",
                is_deleted: false,
                notes: []
            };
            this.info.volunteers.push(newVol);
        },
        deleteEdu: function (i) {
            this.info.educations[i].is_deleted = true;
            // this.$forceUpdate();
        },
        deleteJob: function (i) {
            this.info.employments[i].is_deleted = true
        },
        deleteVol: function (i) {
            this.info.volunteers[i].is_deleted = true
        },
    },

    mounted () {
        axios
            .get('/api/resumejson/' + resumeID)
            .then(response => (this.info = response.data))
    }
});