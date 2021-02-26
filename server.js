const express = require('express')
const bodyParser = require('body-parser')
const jwt = require('jsonwebtoken')
const cors = require('cors')
const app = express()

const secret = "Hello_world"

app.use(bodyParser.json())
app.use(cors())

app.post('/entrar', (req, res) => {
    if (req.body.user == 'admin') {
        if (req.body.password == 'Conviso2021') {
        const token = jwt.sign(secret)
        res.json({ auth: true, token })
        return
        }
    }
    res.status(401).json({ auth: false, message: "Usuário ou senha invalidos." })
})

app.get('/testar', (req, res) => {
    const token = req.headers['authorization'] && req.headers['authorization'].split(' ')[1]
    if (!token) return res.status(401).json({ auth: false, message:"Precisamos de um token."})

    jwt.verify(token, secret, (err, decoded) => {
        if (err) return res.status(500).json({ err, message: "Internal Error"})
        return res.json({ auth: true, msg: "Você passou no teste!" })
    })
})

app.listen(4000, "0.0.0.0", () => {
    console.log("Estamos rodando a API!!!")
})
