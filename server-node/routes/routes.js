const express = require('express');
const router = express.Router();
const userController = require("../controllers/mongo");
const redisController = require("../controllers/redis");

router.get("/", userController.get);
router.get("/top3", userController.getTop3);
router.get("/pie", userController.getPie);
router.get("/bar", redisController.get);
router.get("/last", redisController.getLast);

module.exports = router;