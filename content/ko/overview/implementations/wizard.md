---
title: "매칭 위자드"
description: "요구사항에 맞는 Gateway API 컨트롤러를 찾는 대화형 위자드"
weight: 10
type: docs
---

<div id="wizard-iframe-container">
  <iframe id="wizard-iframe" src="/wizard/?lang=ko" title="Controller matching wizard" scrolling="no" style="width: 100%; border: none; display: block; overflow: hidden;"></iframe>
</div>

<script>
(function() {
  var iframe = document.getElementById('wizard-iframe');
  if (!iframe) return;
  function setHeight(h) {
    iframe.style.height = (typeof h === 'number' ? h : 400) + 'px';
  }
  window.addEventListener('message', function(event) {
    if (event.data && event.data.type === 'wizard-height' && typeof event.data.height === 'number') {
      setHeight(event.data.height);
    }
  });
  setHeight(400);
})();
</script>
