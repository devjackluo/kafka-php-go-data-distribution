<?php
use Monolog\Logger;
use Monolog\Handler\StreamHandler;


if ($_SERVER['REQUEST_METHOD'] === 'POST') {

    $data = file_get_contents('php://input');
    $json = json_decode((string)$data,true);


    $method = $json["endpoint"]["method"];
    $url = $json["endpoint"]["url"];
    $mascot = $json["data"][0]["mascot"];
    $location = $json["data"][0]["location"];

    $url = str_replace("{mascot}",$mascot,$url);
    $url = str_replace("{location}",$location,$url);
    $url = str_replace("{bar}","",$url);


    require './vendor/autoload.php';
    date_default_timezone_set('PRC');


    $logger = new Logger('my_logger');
    $logger->pushHandler(new StreamHandler("php://stdout"));

    $config = \Kafka\ProducerConfig::getInstance();
    $config->setMetadataRefreshIntervalMs(10000);
    $config->setMetadataBrokerList('localhost:9092, localhost:9093, localhost:9094');
    $config->setBrokerVersion('1.0.0');
    $config->setRequiredAck(1);
    $config->setIsAsyn(false);
    $config->setProduceInterval(500);
    $producer = new \Kafka\Producer();

    $producer->send([
        [
            'topic' => 'my_topic',
            'value' => $url,
            'key' => '',
        ],
    ]);

    $producer->setLogger($logger);
    $producer->success(function ($result) {
        var_dump($result);
    });
    $producer->error(function ($errorCode) {
        var_dump($errorCode);
    });
    $producer->send(true);


}else{

    echo "Invalid Request Code Most Likely";

}